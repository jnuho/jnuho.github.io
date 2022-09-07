import boto3
import datetime
from time import sleep
import logging

logger = logging.getLogger()
logger.setLevel(logging.INFO)


"""
######################################################
#
# Lambda Function Invoke Event Parameter JSON Sample
#
######################################################

{
  "AutoScalingGroupName":"AWSDC-LTN-ERP-PRD-RECPRDAP3"
  "InstanceId": "i-00cb4b2aed3c1d77e",
  "AmiName": "AMIAutoScalingTest",
  "AmiDescription": "AMIAutoScalingTest",
  "AmiBillingTag": "KEHRD",
  "LaunchTemplateId": "lt-0528f960dde8a5eb9",
  "LaunchTemplateName":"aws-poc-iis-template",
  "LaunchTemplateInstanceProfileArn":"arn:aws:iam::677963422510:instance-profile/SND_SSM",
  "LaunchInstanceType":"t2.micro",
  "LaunchInstanceSG":["sg-042793ed07d5c7405", "sg-0460c3a0548478d3f"]
}

"""
def lambda_handler(event, context):
 
    logger.info("Start")
   
    AutoScalingGroupName = event["AutoScalingGroupName"] 
    SourceInstanceId = event["InstanceId"]
    AmiName = event["AmiName"]
    AmiDescription = event["AmiDescription"]
    AmiBillingTag = event["AmiBillingTag"]
    LaunchTemplateId = event["LaunchTemplateId"]
    LaunchTemplateName = event["LaunchTemplateName"]
    LaunchTemplateInstanceProfileArn = event["LaunchTemplateInstanceProfileArn"]
    LaunchInstanceType = event["LaunchInstanceType"]
    LaunchInstanceSG = event["LaunchInstanceSG"]
    
    logger.info("Event Parameter : {}".format(event))
    
    ec2Client = boto3.client('ec2')
    
    response = ec2Client.describe_launch_templates(
        LaunchTemplateIds=[LaunchTemplateId]
    )
    
    SourceTemplateVersionNumber = response["LaunchTemplates"][0]["LatestVersionNumber"]
    logger.info("Source LaunchTemplateVersionNumber = {}".format(SourceTemplateVersionNumber))
    
    AmiName = AmiName + "_V" + str(SourceTemplateVersionNumber + 1)
    logger.info("AmiName = {}".format(AmiName))
    
    # Instance AMI 생성
    response = ec2Client.create_image(
        InstanceId=SourceInstanceId
        , Name=AmiName
        , Description=AmiDescription
        , NoReboot=True
    )
    
    AmiId = response["ImageId"]
    
    logger.info("Created AMI [{}]".format(AmiId))
    
    # 생성된 AMI Tagging
    ec2Client.create_tags(
        Resources=[AmiId]
        , Tags=[
              { "Key" : "Name", "Value" : AmiName }
            , { "Key" : "Billing", "Value" : AmiBillingTag }
        ]
    )
    
    logger.info("Start to check AMI state")
    check_ami_state(ec2Client, AmiId)
    logger.info("Ended to check AMI state")
    
    response = ec2Client.create_launch_template_version(
        LaunchTemplateId=LaunchTemplateId
        , SourceVersion=str(SourceTemplateVersionNumber)
        , VersionDescription=LaunchTemplateName
        , LaunchTemplateData={
              "ImageId" : AmiId
            , "InstanceType" : LaunchInstanceType
            , "SecurityGroupIds" : LaunchInstanceSG
            , "IamInstanceProfile" : {
                "Arn" : LaunchTemplateInstanceProfileArn
            }
        }
    )
    
    logger.info("Created LaunchTemplateVersion = {}".format(response))
    
    DefaultVersionNumber = response["LaunchTemplateVersion"]["VersionNumber"]
    
    response = ec2Client.modify_launch_template(
        LaunchTemplateId=LaunchTemplateId
        , DefaultVersion=str(DefaultVersionNumber)
    )
    logger.info("Changed DefaultVersionNumber = {}".format(response))
    
    autoScalingClient = boto3.client('autoscaling')
    
    logger.info(AutoScalingGroupName)
    # 기존 AutoScaling Instances 가져오기
    response = autoScalingClient.describe_auto_scaling_groups(
        AutoScalingGroupNames=[AutoScalingGroupName]
    )
    
    logger.info(response)
    
    # 기존 AutoScaling Instances Terminate
    instances = response["AutoScalingGroups"][0]["Instances"]
    for instance in instances:
        instanceId = instance["InstanceId"]
        response = autoScalingClient.terminate_instance_in_auto_scaling_group(
            InstanceId = instanceId
            , ShouldDecrementDesiredCapacity=False
        )
        logger.info("Instance Terminated - InstanceId = {}".format(instanceId))
        logger.info(response)
        
    logger.info("Ended")
    
    
def check_ami_state(ec2Client, AmiId):
    
    response = ec2Client.describe_images(
        ImageIds=[AmiId]
    )
    
    AmiState = response["Images"][0]["State"]
    logger.info("AmiState = {}".format(AmiState))
    
    if not AmiState == "available":
        sleep(10)
        check_ami_state(ec2Client, AmiId)

