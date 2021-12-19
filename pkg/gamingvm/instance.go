package gamingvm

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/iam"
)

const (
	ownerTagKey = "tag:Owner"
)

var (
	ec2Svc *ec2.EC2
	iamSvc *iam.IAM
)

func init() {
	// Load session from shared config
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create AWS service clients
	ec2Svc = ec2.New(sess)
	iamSvc = iam.New(sess)
}

func GetCurrentUserName() (string, error) {
	if iamSvc == nil {
		return "", fmt.Errorf("iamSvc is nil")
	}

	iamResult, err := iamSvc.GetUser(nil)
	if err != nil {
		return "", err
	}

	return *iamResult.User.UserName, nil
}

func GetGamingVMForUser(username string) (*ec2.Instance, *string, error) {
	// Get instances that we control under our user
	ec2Instances, err := ec2Svc.DescribeInstances(&ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name:   aws.String(ownerTagKey),
				Values: []*string{&username},
			},
		},
	})
	if err != nil {
		return nil, nil, err
	}

	if len(ec2Instances.Reservations) == 0 {
		return nil, nil, fmt.Errorf("found no instances that you (username '%s') can control, check with your administrator", username)
	}

	if len(ec2Instances.Reservations) > 1 {
		return nil, nil, fmt.Errorf("found too many instances (>1) that you (username '%s') can control, check with your administrator", username)
	}

	instance := ec2Instances.Reservations[0].Instances[0]

	var instanceName *string
	for _, tag := range instance.Tags {
		if *tag.Key != "Name" {
			continue
		}

		instanceName = tag.Value
	}

	return instance, instanceName, nil
}

func StartVM(instance *ec2.Instance) error {
	_, err := ec2Svc.StartInstances(&ec2.StartInstancesInput{
		InstanceIds: []*string{instance.InstanceId},
	})
	return err
}

func StopVM(instance *ec2.Instance) error {
	_, err := ec2Svc.StopInstances(&ec2.StopInstancesInput{
		InstanceIds: []*string{instance.InstanceId},
	})
	return err
}
