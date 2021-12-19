# AWS Starter
A simple project I created for a friend of mine who has a gaming VM in my EC2 cloud.
To avoid having them go through the AWS Console or learning to use the command line,
this project creates two simple executables to start and stop the one instance associated
with your user (by user name).

## Configuration
To use this project, the following is required:
* An IAM user with the policy provided below allowing to start and stop certain instances
* An EC2 instance with the tag `Owner=<IAM user name>`
* An AWS config file present with the `[default]` section holding credentials for the IAM user

## Tags
To use this, your EC2 instances are required to have the tag `Owner=${aws.username}`.
E.g. if I have an IAM user `testuser` with the IAM policy provided below, I would add the
label `Owner=testuser` to the instance.

Currently, an error will be raised if multiple instances have this label, or none do.

## Required IAM Policy
I created an IAM policy with minimum required permissions to perform the required operations.

```
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "VisualEditor0",
            "Effect": "Allow",
            "Action": [
                "ec2:StartInstances",
                "ec2:DescribeInstanceAttribute",
                "ec2:StopInstances"
            ],
            "Resource": "arn:aws:ec2:*:*:instance/*",
            "Condition": {
                "StringEquals": {
                    "ec2:ResourceTag/Owner": "${aws:username}"
                }
            }
        },
        {
            "Sid": "VisualEditor1",
            "Effect": "Allow",
            "Action": [
                "ec2:DescribeInstances",
                "ec2:DescribeNetworkInterfaces",
                "ec2:DescribeTags",
                "ec2:DescribeVpcs",
                "ec2:DescribePublicIpv4Pools",
                "ec2:DescribeInstanceTypes",
                "ec2:DescribeInstanceEventWindows",
                "ec2:DescribeInstanceEventNotificationAttributes",
                "ec2:DescribeInstanceStatus"
            ],
            "Resource": "*"
        },
        {
            "Sid": "VisualEditor2",
            "Effect": "Allow",
            "Action": "iam:GetUser",
            "Resource": "arn:aws:iam::*:user/${aws:username}"
        }
    ]
}
```