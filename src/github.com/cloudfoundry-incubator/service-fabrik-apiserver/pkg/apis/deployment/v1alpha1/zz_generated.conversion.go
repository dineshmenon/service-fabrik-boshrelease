// +build !ignore_autogenerated

//TODO copyright header

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v1alpha1

import (
	deployment "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/apis/deployment"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	unsafe "unsafe"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1alpha1_Director_To_deployment_Director,
		Convert_deployment_Director_To_v1alpha1_Director,
		Convert_v1alpha1_DirectorList_To_deployment_DirectorList,
		Convert_deployment_DirectorList_To_v1alpha1_DirectorList,
		Convert_v1alpha1_DirectorSpec_To_deployment_DirectorSpec,
		Convert_deployment_DirectorSpec_To_v1alpha1_DirectorSpec,
		Convert_v1alpha1_DirectorStatus_To_deployment_DirectorStatus,
		Convert_deployment_DirectorStatus_To_v1alpha1_DirectorStatus,
		Convert_v1alpha1_DirectorStatusStrategy_To_deployment_DirectorStatusStrategy,
		Convert_deployment_DirectorStatusStrategy_To_v1alpha1_DirectorStatusStrategy,
		Convert_v1alpha1_DirectorStrategy_To_deployment_DirectorStrategy,
		Convert_deployment_DirectorStrategy_To_v1alpha1_DirectorStrategy,
		Convert_v1alpha1_Docker_To_deployment_Docker,
		Convert_deployment_Docker_To_v1alpha1_Docker,
		Convert_v1alpha1_DockerList_To_deployment_DockerList,
		Convert_deployment_DockerList_To_v1alpha1_DockerList,
		Convert_v1alpha1_DockerSpec_To_deployment_DockerSpec,
		Convert_deployment_DockerSpec_To_v1alpha1_DockerSpec,
		Convert_v1alpha1_DockerStatus_To_deployment_DockerStatus,
		Convert_deployment_DockerStatus_To_v1alpha1_DockerStatus,
		Convert_v1alpha1_DockerStatusStrategy_To_deployment_DockerStatusStrategy,
		Convert_deployment_DockerStatusStrategy_To_v1alpha1_DockerStatusStrategy,
		Convert_v1alpha1_DockerStrategy_To_deployment_DockerStrategy,
		Convert_deployment_DockerStrategy_To_v1alpha1_DockerStrategy,
		Convert_v1alpha1_Virtualhost_To_deployment_Virtualhost,
		Convert_deployment_Virtualhost_To_v1alpha1_Virtualhost,
		Convert_v1alpha1_VirtualhostList_To_deployment_VirtualhostList,
		Convert_deployment_VirtualhostList_To_v1alpha1_VirtualhostList,
		Convert_v1alpha1_VirtualhostSpec_To_deployment_VirtualhostSpec,
		Convert_deployment_VirtualhostSpec_To_v1alpha1_VirtualhostSpec,
		Convert_v1alpha1_VirtualhostStatus_To_deployment_VirtualhostStatus,
		Convert_deployment_VirtualhostStatus_To_v1alpha1_VirtualhostStatus,
		Convert_v1alpha1_VirtualhostStatusStrategy_To_deployment_VirtualhostStatusStrategy,
		Convert_deployment_VirtualhostStatusStrategy_To_v1alpha1_VirtualhostStatusStrategy,
		Convert_v1alpha1_VirtualhostStrategy_To_deployment_VirtualhostStrategy,
		Convert_deployment_VirtualhostStrategy_To_v1alpha1_VirtualhostStrategy,
	)
}

func autoConvert_v1alpha1_Director_To_deployment_Director(in *Director, out *deployment.Director, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_DirectorSpec_To_deployment_DirectorSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_DirectorStatus_To_deployment_DirectorStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_Director_To_deployment_Director is an autogenerated conversion function.
func Convert_v1alpha1_Director_To_deployment_Director(in *Director, out *deployment.Director, s conversion.Scope) error {
	return autoConvert_v1alpha1_Director_To_deployment_Director(in, out, s)
}

func autoConvert_deployment_Director_To_v1alpha1_Director(in *deployment.Director, out *Director, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_deployment_DirectorSpec_To_v1alpha1_DirectorSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_deployment_DirectorStatus_To_v1alpha1_DirectorStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_deployment_Director_To_v1alpha1_Director is an autogenerated conversion function.
func Convert_deployment_Director_To_v1alpha1_Director(in *deployment.Director, out *Director, s conversion.Scope) error {
	return autoConvert_deployment_Director_To_v1alpha1_Director(in, out, s)
}

func autoConvert_v1alpha1_DirectorList_To_deployment_DirectorList(in *DirectorList, out *deployment.DirectorList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]deployment.Director)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_DirectorList_To_deployment_DirectorList is an autogenerated conversion function.
func Convert_v1alpha1_DirectorList_To_deployment_DirectorList(in *DirectorList, out *deployment.DirectorList, s conversion.Scope) error {
	return autoConvert_v1alpha1_DirectorList_To_deployment_DirectorList(in, out, s)
}

func autoConvert_deployment_DirectorList_To_v1alpha1_DirectorList(in *deployment.DirectorList, out *DirectorList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]Director)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_deployment_DirectorList_To_v1alpha1_DirectorList is an autogenerated conversion function.
func Convert_deployment_DirectorList_To_v1alpha1_DirectorList(in *deployment.DirectorList, out *DirectorList, s conversion.Scope) error {
	return autoConvert_deployment_DirectorList_To_v1alpha1_DirectorList(in, out, s)
}

func autoConvert_v1alpha1_DirectorSpec_To_deployment_DirectorSpec(in *DirectorSpec, out *deployment.DirectorSpec, s conversion.Scope) error {
	out.Options = in.Options
	return nil
}

// Convert_v1alpha1_DirectorSpec_To_deployment_DirectorSpec is an autogenerated conversion function.
func Convert_v1alpha1_DirectorSpec_To_deployment_DirectorSpec(in *DirectorSpec, out *deployment.DirectorSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_DirectorSpec_To_deployment_DirectorSpec(in, out, s)
}

func autoConvert_deployment_DirectorSpec_To_v1alpha1_DirectorSpec(in *deployment.DirectorSpec, out *DirectorSpec, s conversion.Scope) error {
	out.Options = in.Options
	return nil
}

// Convert_deployment_DirectorSpec_To_v1alpha1_DirectorSpec is an autogenerated conversion function.
func Convert_deployment_DirectorSpec_To_v1alpha1_DirectorSpec(in *deployment.DirectorSpec, out *DirectorSpec, s conversion.Scope) error {
	return autoConvert_deployment_DirectorSpec_To_v1alpha1_DirectorSpec(in, out, s)
}

func autoConvert_v1alpha1_DirectorStatus_To_deployment_DirectorStatus(in *DirectorStatus, out *deployment.DirectorStatus, s conversion.Scope) error {
	out.State = in.State
	out.LastOperation = in.LastOperation
	out.Response = in.Response
	return nil
}

// Convert_v1alpha1_DirectorStatus_To_deployment_DirectorStatus is an autogenerated conversion function.
func Convert_v1alpha1_DirectorStatus_To_deployment_DirectorStatus(in *DirectorStatus, out *deployment.DirectorStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_DirectorStatus_To_deployment_DirectorStatus(in, out, s)
}

func autoConvert_deployment_DirectorStatus_To_v1alpha1_DirectorStatus(in *deployment.DirectorStatus, out *DirectorStatus, s conversion.Scope) error {
	out.State = in.State
	out.LastOperation = in.LastOperation
	out.Response = in.Response
	return nil
}

// Convert_deployment_DirectorStatus_To_v1alpha1_DirectorStatus is an autogenerated conversion function.
func Convert_deployment_DirectorStatus_To_v1alpha1_DirectorStatus(in *deployment.DirectorStatus, out *DirectorStatus, s conversion.Scope) error {
	return autoConvert_deployment_DirectorStatus_To_v1alpha1_DirectorStatus(in, out, s)
}

func autoConvert_v1alpha1_DirectorStatusStrategy_To_deployment_DirectorStatusStrategy(in *DirectorStatusStrategy, out *deployment.DirectorStatusStrategy, s conversion.Scope) error {
	out.DefaultStatusStorageStrategy = in.DefaultStatusStorageStrategy
	return nil
}

// Convert_v1alpha1_DirectorStatusStrategy_To_deployment_DirectorStatusStrategy is an autogenerated conversion function.
func Convert_v1alpha1_DirectorStatusStrategy_To_deployment_DirectorStatusStrategy(in *DirectorStatusStrategy, out *deployment.DirectorStatusStrategy, s conversion.Scope) error {
	return autoConvert_v1alpha1_DirectorStatusStrategy_To_deployment_DirectorStatusStrategy(in, out, s)
}

func autoConvert_deployment_DirectorStatusStrategy_To_v1alpha1_DirectorStatusStrategy(in *deployment.DirectorStatusStrategy, out *DirectorStatusStrategy, s conversion.Scope) error {
	out.DefaultStatusStorageStrategy = in.DefaultStatusStorageStrategy
	return nil
}

// Convert_deployment_DirectorStatusStrategy_To_v1alpha1_DirectorStatusStrategy is an autogenerated conversion function.
func Convert_deployment_DirectorStatusStrategy_To_v1alpha1_DirectorStatusStrategy(in *deployment.DirectorStatusStrategy, out *DirectorStatusStrategy, s conversion.Scope) error {
	return autoConvert_deployment_DirectorStatusStrategy_To_v1alpha1_DirectorStatusStrategy(in, out, s)
}

func autoConvert_v1alpha1_DirectorStrategy_To_deployment_DirectorStrategy(in *DirectorStrategy, out *deployment.DirectorStrategy, s conversion.Scope) error {
	out.DefaultStorageStrategy = in.DefaultStorageStrategy
	return nil
}

// Convert_v1alpha1_DirectorStrategy_To_deployment_DirectorStrategy is an autogenerated conversion function.
func Convert_v1alpha1_DirectorStrategy_To_deployment_DirectorStrategy(in *DirectorStrategy, out *deployment.DirectorStrategy, s conversion.Scope) error {
	return autoConvert_v1alpha1_DirectorStrategy_To_deployment_DirectorStrategy(in, out, s)
}

func autoConvert_deployment_DirectorStrategy_To_v1alpha1_DirectorStrategy(in *deployment.DirectorStrategy, out *DirectorStrategy, s conversion.Scope) error {
	out.DefaultStorageStrategy = in.DefaultStorageStrategy
	return nil
}

// Convert_deployment_DirectorStrategy_To_v1alpha1_DirectorStrategy is an autogenerated conversion function.
func Convert_deployment_DirectorStrategy_To_v1alpha1_DirectorStrategy(in *deployment.DirectorStrategy, out *DirectorStrategy, s conversion.Scope) error {
	return autoConvert_deployment_DirectorStrategy_To_v1alpha1_DirectorStrategy(in, out, s)
}

func autoConvert_v1alpha1_Docker_To_deployment_Docker(in *Docker, out *deployment.Docker, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_DockerSpec_To_deployment_DockerSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_DockerStatus_To_deployment_DockerStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_Docker_To_deployment_Docker is an autogenerated conversion function.
func Convert_v1alpha1_Docker_To_deployment_Docker(in *Docker, out *deployment.Docker, s conversion.Scope) error {
	return autoConvert_v1alpha1_Docker_To_deployment_Docker(in, out, s)
}

func autoConvert_deployment_Docker_To_v1alpha1_Docker(in *deployment.Docker, out *Docker, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_deployment_DockerSpec_To_v1alpha1_DockerSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_deployment_DockerStatus_To_v1alpha1_DockerStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_deployment_Docker_To_v1alpha1_Docker is an autogenerated conversion function.
func Convert_deployment_Docker_To_v1alpha1_Docker(in *deployment.Docker, out *Docker, s conversion.Scope) error {
	return autoConvert_deployment_Docker_To_v1alpha1_Docker(in, out, s)
}

func autoConvert_v1alpha1_DockerList_To_deployment_DockerList(in *DockerList, out *deployment.DockerList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]deployment.Docker)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_DockerList_To_deployment_DockerList is an autogenerated conversion function.
func Convert_v1alpha1_DockerList_To_deployment_DockerList(in *DockerList, out *deployment.DockerList, s conversion.Scope) error {
	return autoConvert_v1alpha1_DockerList_To_deployment_DockerList(in, out, s)
}

func autoConvert_deployment_DockerList_To_v1alpha1_DockerList(in *deployment.DockerList, out *DockerList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]Docker)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_deployment_DockerList_To_v1alpha1_DockerList is an autogenerated conversion function.
func Convert_deployment_DockerList_To_v1alpha1_DockerList(in *deployment.DockerList, out *DockerList, s conversion.Scope) error {
	return autoConvert_deployment_DockerList_To_v1alpha1_DockerList(in, out, s)
}

func autoConvert_v1alpha1_DockerSpec_To_deployment_DockerSpec(in *DockerSpec, out *deployment.DockerSpec, s conversion.Scope) error {
	out.Options = in.Options
	return nil
}

// Convert_v1alpha1_DockerSpec_To_deployment_DockerSpec is an autogenerated conversion function.
func Convert_v1alpha1_DockerSpec_To_deployment_DockerSpec(in *DockerSpec, out *deployment.DockerSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_DockerSpec_To_deployment_DockerSpec(in, out, s)
}

func autoConvert_deployment_DockerSpec_To_v1alpha1_DockerSpec(in *deployment.DockerSpec, out *DockerSpec, s conversion.Scope) error {
	out.Options = in.Options
	return nil
}

// Convert_deployment_DockerSpec_To_v1alpha1_DockerSpec is an autogenerated conversion function.
func Convert_deployment_DockerSpec_To_v1alpha1_DockerSpec(in *deployment.DockerSpec, out *DockerSpec, s conversion.Scope) error {
	return autoConvert_deployment_DockerSpec_To_v1alpha1_DockerSpec(in, out, s)
}

func autoConvert_v1alpha1_DockerStatus_To_deployment_DockerStatus(in *DockerStatus, out *deployment.DockerStatus, s conversion.Scope) error {
	out.State = in.State
	out.LastOperation = in.LastOperation
	out.Response = in.Response
	return nil
}

// Convert_v1alpha1_DockerStatus_To_deployment_DockerStatus is an autogenerated conversion function.
func Convert_v1alpha1_DockerStatus_To_deployment_DockerStatus(in *DockerStatus, out *deployment.DockerStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_DockerStatus_To_deployment_DockerStatus(in, out, s)
}

func autoConvert_deployment_DockerStatus_To_v1alpha1_DockerStatus(in *deployment.DockerStatus, out *DockerStatus, s conversion.Scope) error {
	out.State = in.State
	out.LastOperation = in.LastOperation
	out.Response = in.Response
	return nil
}

// Convert_deployment_DockerStatus_To_v1alpha1_DockerStatus is an autogenerated conversion function.
func Convert_deployment_DockerStatus_To_v1alpha1_DockerStatus(in *deployment.DockerStatus, out *DockerStatus, s conversion.Scope) error {
	return autoConvert_deployment_DockerStatus_To_v1alpha1_DockerStatus(in, out, s)
}

func autoConvert_v1alpha1_DockerStatusStrategy_To_deployment_DockerStatusStrategy(in *DockerStatusStrategy, out *deployment.DockerStatusStrategy, s conversion.Scope) error {
	out.DefaultStatusStorageStrategy = in.DefaultStatusStorageStrategy
	return nil
}

// Convert_v1alpha1_DockerStatusStrategy_To_deployment_DockerStatusStrategy is an autogenerated conversion function.
func Convert_v1alpha1_DockerStatusStrategy_To_deployment_DockerStatusStrategy(in *DockerStatusStrategy, out *deployment.DockerStatusStrategy, s conversion.Scope) error {
	return autoConvert_v1alpha1_DockerStatusStrategy_To_deployment_DockerStatusStrategy(in, out, s)
}

func autoConvert_deployment_DockerStatusStrategy_To_v1alpha1_DockerStatusStrategy(in *deployment.DockerStatusStrategy, out *DockerStatusStrategy, s conversion.Scope) error {
	out.DefaultStatusStorageStrategy = in.DefaultStatusStorageStrategy
	return nil
}

// Convert_deployment_DockerStatusStrategy_To_v1alpha1_DockerStatusStrategy is an autogenerated conversion function.
func Convert_deployment_DockerStatusStrategy_To_v1alpha1_DockerStatusStrategy(in *deployment.DockerStatusStrategy, out *DockerStatusStrategy, s conversion.Scope) error {
	return autoConvert_deployment_DockerStatusStrategy_To_v1alpha1_DockerStatusStrategy(in, out, s)
}

func autoConvert_v1alpha1_DockerStrategy_To_deployment_DockerStrategy(in *DockerStrategy, out *deployment.DockerStrategy, s conversion.Scope) error {
	out.DefaultStorageStrategy = in.DefaultStorageStrategy
	return nil
}

// Convert_v1alpha1_DockerStrategy_To_deployment_DockerStrategy is an autogenerated conversion function.
func Convert_v1alpha1_DockerStrategy_To_deployment_DockerStrategy(in *DockerStrategy, out *deployment.DockerStrategy, s conversion.Scope) error {
	return autoConvert_v1alpha1_DockerStrategy_To_deployment_DockerStrategy(in, out, s)
}

func autoConvert_deployment_DockerStrategy_To_v1alpha1_DockerStrategy(in *deployment.DockerStrategy, out *DockerStrategy, s conversion.Scope) error {
	out.DefaultStorageStrategy = in.DefaultStorageStrategy
	return nil
}

// Convert_deployment_DockerStrategy_To_v1alpha1_DockerStrategy is an autogenerated conversion function.
func Convert_deployment_DockerStrategy_To_v1alpha1_DockerStrategy(in *deployment.DockerStrategy, out *DockerStrategy, s conversion.Scope) error {
	return autoConvert_deployment_DockerStrategy_To_v1alpha1_DockerStrategy(in, out, s)
}

func autoConvert_v1alpha1_Virtualhost_To_deployment_Virtualhost(in *Virtualhost, out *deployment.Virtualhost, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_VirtualhostSpec_To_deployment_VirtualhostSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_VirtualhostStatus_To_deployment_VirtualhostStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_Virtualhost_To_deployment_Virtualhost is an autogenerated conversion function.
func Convert_v1alpha1_Virtualhost_To_deployment_Virtualhost(in *Virtualhost, out *deployment.Virtualhost, s conversion.Scope) error {
	return autoConvert_v1alpha1_Virtualhost_To_deployment_Virtualhost(in, out, s)
}

func autoConvert_deployment_Virtualhost_To_v1alpha1_Virtualhost(in *deployment.Virtualhost, out *Virtualhost, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_deployment_VirtualhostSpec_To_v1alpha1_VirtualhostSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_deployment_VirtualhostStatus_To_v1alpha1_VirtualhostStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_deployment_Virtualhost_To_v1alpha1_Virtualhost is an autogenerated conversion function.
func Convert_deployment_Virtualhost_To_v1alpha1_Virtualhost(in *deployment.Virtualhost, out *Virtualhost, s conversion.Scope) error {
	return autoConvert_deployment_Virtualhost_To_v1alpha1_Virtualhost(in, out, s)
}

func autoConvert_v1alpha1_VirtualhostList_To_deployment_VirtualhostList(in *VirtualhostList, out *deployment.VirtualhostList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]deployment.Virtualhost)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_VirtualhostList_To_deployment_VirtualhostList is an autogenerated conversion function.
func Convert_v1alpha1_VirtualhostList_To_deployment_VirtualhostList(in *VirtualhostList, out *deployment.VirtualhostList, s conversion.Scope) error {
	return autoConvert_v1alpha1_VirtualhostList_To_deployment_VirtualhostList(in, out, s)
}

func autoConvert_deployment_VirtualhostList_To_v1alpha1_VirtualhostList(in *deployment.VirtualhostList, out *VirtualhostList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]Virtualhost)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_deployment_VirtualhostList_To_v1alpha1_VirtualhostList is an autogenerated conversion function.
func Convert_deployment_VirtualhostList_To_v1alpha1_VirtualhostList(in *deployment.VirtualhostList, out *VirtualhostList, s conversion.Scope) error {
	return autoConvert_deployment_VirtualhostList_To_v1alpha1_VirtualhostList(in, out, s)
}

func autoConvert_v1alpha1_VirtualhostSpec_To_deployment_VirtualhostSpec(in *VirtualhostSpec, out *deployment.VirtualhostSpec, s conversion.Scope) error {
	out.Options = in.Options
	return nil
}

// Convert_v1alpha1_VirtualhostSpec_To_deployment_VirtualhostSpec is an autogenerated conversion function.
func Convert_v1alpha1_VirtualhostSpec_To_deployment_VirtualhostSpec(in *VirtualhostSpec, out *deployment.VirtualhostSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_VirtualhostSpec_To_deployment_VirtualhostSpec(in, out, s)
}

func autoConvert_deployment_VirtualhostSpec_To_v1alpha1_VirtualhostSpec(in *deployment.VirtualhostSpec, out *VirtualhostSpec, s conversion.Scope) error {
	out.Options = in.Options
	return nil
}

// Convert_deployment_VirtualhostSpec_To_v1alpha1_VirtualhostSpec is an autogenerated conversion function.
func Convert_deployment_VirtualhostSpec_To_v1alpha1_VirtualhostSpec(in *deployment.VirtualhostSpec, out *VirtualhostSpec, s conversion.Scope) error {
	return autoConvert_deployment_VirtualhostSpec_To_v1alpha1_VirtualhostSpec(in, out, s)
}

func autoConvert_v1alpha1_VirtualhostStatus_To_deployment_VirtualhostStatus(in *VirtualhostStatus, out *deployment.VirtualhostStatus, s conversion.Scope) error {
	out.State = in.State
	out.Error = in.Error
	out.LastOperation = in.LastOperation
	out.Response = in.Response
	return nil
}

// Convert_v1alpha1_VirtualhostStatus_To_deployment_VirtualhostStatus is an autogenerated conversion function.
func Convert_v1alpha1_VirtualhostStatus_To_deployment_VirtualhostStatus(in *VirtualhostStatus, out *deployment.VirtualhostStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_VirtualhostStatus_To_deployment_VirtualhostStatus(in, out, s)
}

func autoConvert_deployment_VirtualhostStatus_To_v1alpha1_VirtualhostStatus(in *deployment.VirtualhostStatus, out *VirtualhostStatus, s conversion.Scope) error {
	out.State = in.State
	out.Error = in.Error
	out.LastOperation = in.LastOperation
	out.Response = in.Response
	return nil
}

// Convert_deployment_VirtualhostStatus_To_v1alpha1_VirtualhostStatus is an autogenerated conversion function.
func Convert_deployment_VirtualhostStatus_To_v1alpha1_VirtualhostStatus(in *deployment.VirtualhostStatus, out *VirtualhostStatus, s conversion.Scope) error {
	return autoConvert_deployment_VirtualhostStatus_To_v1alpha1_VirtualhostStatus(in, out, s)
}

func autoConvert_v1alpha1_VirtualhostStatusStrategy_To_deployment_VirtualhostStatusStrategy(in *VirtualhostStatusStrategy, out *deployment.VirtualhostStatusStrategy, s conversion.Scope) error {
	out.DefaultStatusStorageStrategy = in.DefaultStatusStorageStrategy
	return nil
}

// Convert_v1alpha1_VirtualhostStatusStrategy_To_deployment_VirtualhostStatusStrategy is an autogenerated conversion function.
func Convert_v1alpha1_VirtualhostStatusStrategy_To_deployment_VirtualhostStatusStrategy(in *VirtualhostStatusStrategy, out *deployment.VirtualhostStatusStrategy, s conversion.Scope) error {
	return autoConvert_v1alpha1_VirtualhostStatusStrategy_To_deployment_VirtualhostStatusStrategy(in, out, s)
}

func autoConvert_deployment_VirtualhostStatusStrategy_To_v1alpha1_VirtualhostStatusStrategy(in *deployment.VirtualhostStatusStrategy, out *VirtualhostStatusStrategy, s conversion.Scope) error {
	out.DefaultStatusStorageStrategy = in.DefaultStatusStorageStrategy
	return nil
}

// Convert_deployment_VirtualhostStatusStrategy_To_v1alpha1_VirtualhostStatusStrategy is an autogenerated conversion function.
func Convert_deployment_VirtualhostStatusStrategy_To_v1alpha1_VirtualhostStatusStrategy(in *deployment.VirtualhostStatusStrategy, out *VirtualhostStatusStrategy, s conversion.Scope) error {
	return autoConvert_deployment_VirtualhostStatusStrategy_To_v1alpha1_VirtualhostStatusStrategy(in, out, s)
}

func autoConvert_v1alpha1_VirtualhostStrategy_To_deployment_VirtualhostStrategy(in *VirtualhostStrategy, out *deployment.VirtualhostStrategy, s conversion.Scope) error {
	out.DefaultStorageStrategy = in.DefaultStorageStrategy
	return nil
}

// Convert_v1alpha1_VirtualhostStrategy_To_deployment_VirtualhostStrategy is an autogenerated conversion function.
func Convert_v1alpha1_VirtualhostStrategy_To_deployment_VirtualhostStrategy(in *VirtualhostStrategy, out *deployment.VirtualhostStrategy, s conversion.Scope) error {
	return autoConvert_v1alpha1_VirtualhostStrategy_To_deployment_VirtualhostStrategy(in, out, s)
}

func autoConvert_deployment_VirtualhostStrategy_To_v1alpha1_VirtualhostStrategy(in *deployment.VirtualhostStrategy, out *VirtualhostStrategy, s conversion.Scope) error {
	out.DefaultStorageStrategy = in.DefaultStorageStrategy
	return nil
}

// Convert_deployment_VirtualhostStrategy_To_v1alpha1_VirtualhostStrategy is an autogenerated conversion function.
func Convert_deployment_VirtualhostStrategy_To_v1alpha1_VirtualhostStrategy(in *deployment.VirtualhostStrategy, out *VirtualhostStrategy, s conversion.Scope) error {
	return autoConvert_deployment_VirtualhostStrategy_To_v1alpha1_VirtualhostStrategy(in, out, s)
}
