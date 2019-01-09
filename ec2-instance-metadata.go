package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const metadataURL = "http://169.254.169.254/latest/meta-data/"

func printValue(resourcePath string) {
	if resourcePath == "options" {
		resourcePath = ""
	}

	response, err := http.Get(metadataURL + resourcePath)

	if err != nil {
		fmt.Printf("Failed to retrieve %s, cause: %t\n", resourcePath, err)
		return
	}

	bodyBytes, _ := ioutil.ReadAll(response.Body)
	bodyString := string(bodyBytes)

	if resourcePath == "options" {
		fmt.Printf("%s\n\n", bodyString)
	} else {
		fmt.Printf("%s: %s\n\n", resourcePath, bodyString)
	}
}

func printHelp() {
	fmt.Println("Welcome to EC2 meta-data retriever")
	fmt.Println("Usage:")
	fmt.Println("$ ec2-metadata <option>")
	fmt.Println("")
	fmt.Println("These are some of the valid options, for a complete list use the option: options")
	fmt.Println("")
	fmt.Println(" - ami-id - The AMI ID used to launch the instance.")
	fmt.Println(" - ami-launch-index - If you started more than one instance at the same time, this value indicates the order in which the instance was launched. The value of the first instance launched is 0.")
	fmt.Println(" - ami-manifest-path - The path to the AMI manifest file in Amazon S3. If you used an Amazon EBS-backed AMI to launch the instance, the returned result is unknown.")
	fmt.Println(" - ancestor-ami-ids - The AMI IDs of any instances that were rebundled to create this AMI. This value will only exist if the AMI manifest file contained an ancestor-amis key.")
	fmt.Println(" - block-device-mapping/ami - The virtual device that contains the root/boot file system.")
	fmt.Println(" - block-device-mapping/ebsN - The virtual devices associated with Amazon EBS volumes, if any are present. Amazon EBS volumes are only available in metadata if they were present at launch time or when the instance was last started. The N indicates the index of the Amazon EBS volume (such as ebs1 or ebs2).")
	fmt.Println(" - block-device-mapping/ephemeralN - The virtual devices associated with non-NVMe instance store volumes, if any are present. The N indicates the index of each ephemeral volume.")
	fmt.Println(" - block-device-mapping/root - The virtual devices or partitions associated with the root devices, or partitions on the virtual device, where the root (/ or C:) file system is associated with the given instance.")
	fmt.Println(" - block-device-mapping/swap - The virtual devices associated with swap. Not always present.")
	fmt.Println(" - hostname - The private IPv4 DNS hostname of the instance. In cases where multiple network interfaces are present, this refers to the eth0 device (the device for which the device number is 0).")
	fmt.Println(" - iam/info - If there is an IAM role associated with the instance, contains information about the last time the instance profile was updated, including the instance's LastUpdated date, InstanceProfileArn, and InstanceProfileId. Otherwise, not present.")
	fmt.Println(" - instance-id - The ID of this instance.")
	fmt.Println(" - instance-type - The type of instance.")
	fmt.Println(" - kernel-id - The ID of the kernel launched with this instance, if applicable.")
	fmt.Println(" - local-hostname - The private IPv4 DNS hostname of the instance. In cases where multiple network interfaces are present, this refers to the eth0 device (the device for which the device number is 0).")
	fmt.Println(" - local-ipv4 - The private IPv4 address of the instance. In cases where multiple network interfaces are present, this refers to the eth0 device (the device for which the device number is 0). ")
	fmt.Println(" - mac - The instance's media access control (MAC) address. In cases where multiple network interfaces are present, this refers to the eth0 device (the device for which the device number is 0).")
	fmt.Println(" - product-codes - Marketplace product codes associated with the instance, if any.")
	fmt.Println(" - public-hostname - The instance's public DNS. This category is only returned if the enableDnsHostnames attribute is set to true.")
	fmt.Println(" - public-ipv4 - The public IPv4 address. If an Elastic IP address is associated with the instance, the value returned is the Elastic IP address.")
	fmt.Println(" - public-keys/0/openssh-key - Public key. Only available if supplied at instance launch time.")
	fmt.Println(" - ramdisk-id - The ID of the RAM disk specified at launch time, if applicable.")
	fmt.Println(" - reservation-id - The ID of the reservation.")
	fmt.Println(" - security-groups - The names of the security groups applied to the instance.")
	fmt.Println(" - services/domain - The domain for AWS resources for the region.")
	fmt.Println(" - services/partition - The partition that the resource is in. For standard AWS regions, the partition is aws. If you have resources in other partitions, the partition is aws-partitionname. For example, the partition for resources in the China (Beijing) region is aws-cn.")
}

func main() {
	if len(os.Args) == 1 {
		printHelp()
		os.Exit(0)
	}

	resource := os.Args[1]
	resource = strings.ToLower(resource)

	printValue(resource)
}
