// Copyright 2021-2023 vArmor Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package apparmor

import (
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/jinzhu/copier"
	kubearmorenforcer "github.com/kubearmor/KubeArmor/KubeArmor/enforcer"
	kubearmortype "github.com/kubearmor/KubeArmor/KubeArmor/types"

	varmor "github.com/bytedance/vArmor/apis/varmor/v1beta1"
)

func GenerateAlwaysAllowProfile(profileName string) string {
	c := []byte(fmt.Sprintf(alwaysAllowTemplate, profileName, ""))
	return base64.StdEncoding.EncodeToString(c)
}

func GenerateRuntimeDefaultProfile(profileName string) string {
	c := []byte(fmt.Sprintf(runtimeDefaultTemplate, profileName, profileName, profileName, ""))
	return base64.StdEncoding.EncodeToString(c)
}

func generateHardeningRules(rule string) (rules string) {
	switch strings.ToLower(rule) {
	//// 1. Blocking escape vectors from privileged container
	// disallow write core_pattern
	case "disallow_write_core_pattern":
		rules += "  deny /proc/sys/kernel/core_pattern w,\n"
	// disallow mount procfs
	case "disallow_mount_procfs":
		rules += "  deny mount fstype=proc,\n"
	// disallow write release_agent
	case "disallow_write_release_agent":
		rules += "  deny /sys/fs/cgroup/**/release_agent w,\n"
	// disallow mount cgroupfs
	case "disallow_mount_cgroupfs":
		rules += "  deny mount fstype=cgroup,\n"
	// disallow debug disk devices
	case "disallow_debug_disk_device":
		rules += "{{range $value := .DiskDevices}}"
		rules += "  deny /dev/{{$value}} rw,\n"
		rules += "{{end}}"
	// disallow mount disk devices
	case "disallow_mount_disk_device":
		rules += "{{range $value := .DiskDevices}}"
		rules += "  deny mount /dev/{{$value}},\n"
		rules += "{{end}}"
	// disallow mount anything
	case "disallow_mount":
		rules += "  deny mount,\n"
	// disallow insmond
	case "disallow_insmod":
		rules += "  deny capability sys_module,\n"
	// disallow load ebpf program
	case "disallow_load_ebpf":
		rules += "  deny capability sys_admin,\n"
		rules += "  deny capability bpf,\n"
	// disallow access to the root of the task through procfs
	case "disallow_access_procfs_root":
		rules += "  deny ptrace read,\n"

	//// 2. Disable capabilities
	// disable all capabilities
	case "disable_cap_all":
		rules += "  deny capability,\n"
	// disable privileged capabilities
	case "disable_cap_privileged":
		rules += `  deny capability dac_read_search,
  deny capability linux_immutable,
  deny capability net_broadcast,
  deny capability net_admin,
  deny capability ipc_lock,
  deny capability ipc_owner,
  deny capability sys_module,
  deny capability sys_rawio,
  deny capability sys_ptrace,
  deny capability sys_pacct,
  deny capability sys_admin,
  deny capability sys_boot,
  deny capability sys_nice,
  deny capability sys_resource,
  deny capability sys_time,
  deny capability sys_tty_config,
  deny capability lease,
  deny capability audit_control,
  deny capability mac_override,
  deny capability mac_admin,
  deny capability syslog,
  deny capability wake_alarm,
  deny capability block_suspend,
  deny capability audit_read,
`
	// disable the specified capability
	case "disable_cap_chown":
		rules += "  deny capability chown,\n"
	case "disable_cap_dac_override":
		rules += "  deny capability dac_override,\n"
	case "disable_cap_dac_read_search":
		rules += "  deny capability dac_read_search,\n"
	case "disable_cap_fowner":
		rules += "  deny capability fowner,\n"
	case "disable_cap_fsetid":
		rules += "  deny capability fsetid,\n"
	case "disable_cap_kill":
		rules += "  deny capability kill,\n"
	case "disable_cap_setgid":
		rules += "  deny capability setgid,\n"
	case "disable_cap_setuid":
		rules += "  deny capability setuid,\n"
	case "disable_cap_setpcap":
		rules += "  deny capability setpcap,\n"
	case "disable_cap_linux_immutable":
		rules += "  deny capability linux_immutable,\n"
	case "disable_cap_net_bind_service":
		rules += "  deny capability net_bind_service,\n"
	case "disable_cap_net_broadcast":
		rules += "  deny capability net_broadcast,\n"
	case "disable_cap_net_admin":
		rules += "  deny capability net_admin,\n"
	case "disable_cap_net_raw":
		rules += "  deny capability net_raw,\n"
	case "disable_cap_ipc_lock":
		rules += "  deny capability ipc_lock,\n"
	case "disable_cap_ipc_owner":
		rules += "  deny capability ipc_owner,\n"
	case "disable_cap_sys_module":
		rules += "  deny capability sys_module,\n"
	case "disable_cap_sys_rawio":
		rules += "  deny capability sys_rawio,\n"
	case "disable_cap_sys_chroot":
		rules += "  deny capability sys_chroot,\n"
	case "disable_cap_sys_ptrace":
		rules += "  deny capability sys_ptrace,\n"
	case "disable_cap_sys_pacct":
		rules += "  deny capability sys_pacct,\n"
	case "disable_cap_sys_admin":
		rules += "  deny capability sys_admin,\n"
	case "disable_cap_sys_boot":
		rules += "  deny capability sys_boot,\n"
	case "disable_cap_sys_nice":
		rules += "  deny capability sys_nice,\n"
	case "disable_cap_sys_resource":
		rules += "  deny capability sys_resource,\n"
	case "disable_cap_sys_time":
		rules += "  deny capability sys_time,\n"
	case "disable_cap_sys_tty_config":
		rules += "  deny capability sys_tty_config,\n"
	case "disable_cap_mknod":
		rules += "  deny capability mknod,\n"
	case "disable_cap_lease":
		rules += "  deny capability lease,\n"
	case "disable_cap_audit_write":
		rules += "  deny capability audit_write,\n"
	case "disable_cap_audit_control":
		rules += "  deny capability audit_control,\n"
	case "disable_cap_setfcap":
		rules += "  deny capability setfcap,\n"
	case "disable_cap_mac_override":
		rules += "  deny capability mac_override,\n"
	case "disable_cap_mac_admin":
		rules += "  deny capability mac_admin,\n"
	case "disable_cap_syslog":
		rules += "  deny capability syslog,\n"
	case "disable_cap_wake_alarm":
		rules += "  deny capability wake_alarm,\n"
	case "disable_cap_block_suspend":
		rules += "  deny capability block_suspend,\n"
	case "disable_cap_audit_read":
		rules += "  deny capability audit_read,\n"
	case "disable_cap_perfmon":
		rules += "  deny capability perfmon,\n"
	case "disable_cap_bpf":
		rules += "  deny capability bpf,\n"

	//// 3. Kernel vulnerability mitigation
	// forward-compatible
	case "disallow_create_user_ns":
		fallthrough
	// diallow abuse user namespace
	case "disallow_abuse_user_ns":
		rules += "  deny capability sys_admin,\n"
	}
	return rules
}

func generateAttackProtectionRules(rule string) (rules string) {
	switch strings.ToLower(rule) {
	//// 4. Mitigate container information leakage
	case "mitigate-sa-leak":
		rules += "  deny /run/secrets/kubernetes.io/serviceaccount/** r,\n"
		rules += "  deny /var/run/secrets/kubernetes.io/serviceaccount/** r,\n"
	case "mitigate-disk-device-number-leak":
		rules += "  deny /proc/partitions r,\n"
		rules += "  deny /proc/**/mountinfo r,\n"
	case "mitigate-overlayfs-leak":
		rules += "  deny /proc/**/mounts r,\n"
		rules += "  deny /proc/**/mountinfo r,\n"
	case "mitigate-host-ip-leak":
		rules += "  deny /proc/**/net/arp r,\n"
	//// 5. Restrict the execution of sensitive commands inside the container
	case "disable-write-etc":
		rules += "  deny /etc/** wl,\n"
	case "disable-busybox":
		rules += "  deny /**/busybox rx,\n"
	case "disable-shell":
		rules += "  deny /**/sh rx,\n"
		rules += "  deny /**/bash rx,\n"
		rules += "  deny /**/dash rx,\n"
	case "disable-wget":
		rules += "  deny /**/wget rx,\n"
	case "disable-curl":
		rules += "  deny /**/curl rx,\n"
	case "disable-chmod":
		rules += "  deny /**/chmod rx,\n"
	case "disable-su-sudo":
		rules += "  deny /**/su rx,\n"
		rules += "  deny /**/sudo rx,\n"
	}
	return rules
}

func generateVulMitigationRules(rule string) (rules string) {
	switch strings.ToLower(rule) {
	case "cgroups_lxcfs_escape_mitigation":
		rules += "  deny /**/release_agent w,\n"
		rules += "  deny /**/devices/devices.allow w,\n"
		rules += "  deny /**/devices/**/devices.allow w,\n"
		rules += "  deny /**/devices/cgroup.procs w,\n"
		rules += "  deny /**/devices/**/cgroup.procs w,\n"
		rules += "  deny /**/devices/tasks w,\n"
		rules += "  deny /**/devices/**/tasks w,\n"
	}
	return rules
}

func GenerateEnhanceProtectProfile(enhanceProtect *varmor.EnhanceProtect, profileName string, privileged bool) string {
	var baseRules string

	// Hardening
	for _, rule := range enhanceProtect.HardeningRules {
		baseRules += generateHardeningRules(rule)
	}

	// Vulnerability Mitigation
	for _, rule := range enhanceProtect.VulMitigationRules {
		baseRules += generateVulMitigationRules(rule)
	}

	// Custom
	for _, rule := range enhanceProtect.AppArmorRawRules {
		if strings.HasSuffix(rule, ",") {
			baseRules += "  " + rule + "\n"
		}
	}

	// Attack Protection
	for _, attackProtectionRule := range enhanceProtect.AttackProtectionRules {
		if len(attackProtectionRule.Targets) == 0 {
			for _, rule := range attackProtectionRule.Rules {
				baseRules += generateAttackProtectionRules(rule)
			}
		}
	}

	// childName(target): childRules
	childRulesMap := make(map[string]string)
	for _, attackProtectionRule := range enhanceProtect.AttackProtectionRules {
		if len(attackProtectionRule.Targets) != 0 {
			var childRules string

			for _, childName := range attackProtectionRule.Targets {
				if _, ok := childRulesMap[childName]; !ok {
					childRules = baseRules
				} else {
					childRules = childRulesMap[childName]
				}

				for _, rule := range attackProtectionRule.Rules {
					childRules += generateAttackProtectionRules(rule)
					childRulesMap[childName] = childRules
				}
			}
		}
	}

	for childName, childRules := range childRulesMap {
		if privileged {
			// Create the child profile for privileged container based on the AlwaysAllow child template
			baseRules += fmt.Sprintf(alwaysAllowChildTemplate, childName, childName, childName, childRules)
		} else {
			// Create the child profile for unprivileged container based on the RuntimeDefault child template
			childProfileName := fmt.Sprintf("%s//%s", profileName, childName)
			baseRules += fmt.Sprintf(runtimeDefaultChildTemplate,
				childProfileName,                // signal
				childName, childName, childName, // target
				profileName, childProfileName, // signal
				profileName, childProfileName, // ptrace
				childRules)
		}
	}

	if privileged {
		// Create profile for privileged container based on the AlwaysAllow template
		p := fmt.Sprintf(alwaysAllowTemplate, profileName, baseRules)
		return base64.StdEncoding.EncodeToString([]byte(p))
	} else {
		// Create profile for unprivileged container based on the RuntimeDefault template
		p := fmt.Sprintf(runtimeDefaultTemplate, profileName, profileName, profileName, baseRules)
		return base64.StdEncoding.EncodeToString([]byte(p))
	}
}

// See WatchHostSecurityPolicies in KubeArmor/core/kubeUpdate.go
func initCustomPolicy(policy *varmor.CustomPolicy) {
	// TODO: Add some validation here for varmor.CustomPolicy.

	for i, e := range policy.Process.MatchPaths {
		if e.Action == "" {
			if policy.Process.Action == "" {
				policy.Process.MatchPaths[i].Action = policy.Action
			} else {
				policy.Process.MatchPaths[i].Action = policy.Process.Action
			}
		}
	}

	for i, e := range policy.Process.MatchDirectories {
		if e.Action == "" {
			if policy.Process.Action == "" {
				policy.Process.MatchDirectories[i].Action = policy.Action
			} else {
				policy.Process.MatchDirectories[i].Action = policy.Process.Action
			}
		}
	}

	for i, e := range policy.Process.MatchPatterns {
		if e.Action == "" {
			if policy.Process.Action == "" {
				policy.Process.MatchPatterns[i].Action = policy.Action
			} else {
				policy.Process.MatchPatterns[i].Action = policy.Process.Action
			}
		}
	}

	for i, e := range policy.File.MatchPaths {
		if e.Action == "" {
			if policy.File.Action == "" {
				policy.File.MatchPaths[i].Action = policy.Action
			} else {
				policy.File.MatchPaths[i].Action = policy.File.Action
			}
		}
	}

	for i, e := range policy.File.MatchDirectories {
		if e.Action == "" {
			if policy.File.Action == "" {
				policy.File.MatchDirectories[i].Action = policy.Action
			} else {
				policy.File.MatchDirectories[i].Action = policy.File.Action
			}
		}
	}

	for i, e := range policy.File.MatchPatterns {
		if e.Action == "" {
			if policy.File.Action == "" {
				policy.File.MatchPatterns[i].Action = policy.Action
			} else {
				policy.File.MatchPatterns[i].Action = policy.File.Action
			}
		}
	}

	for i, e := range policy.Network.MatchProtocols {
		if e.Action == "" {
			if policy.Network.Action == "" {
				policy.Network.MatchProtocols[i].Action = policy.Action
			} else {
				policy.Network.MatchProtocols[i].Action = policy.Network.Action
			}
		}
	}

	for i, e := range policy.Capabilities.MatchCapabilities {
		if e.Action == "" {
			if policy.Capabilities.Action == "" {
				policy.Capabilities.MatchCapabilities[i].Action = policy.Action
			} else {
				policy.Capabilities.MatchCapabilities[i].Action = policy.Capabilities.Action
			}
		}
	}
}

func newKubeArmorSecurityPolicy(policy varmor.CustomPolicy) *kubearmortype.SecurityPolicy {
	customPolicy := kubearmortype.SecurityPolicy{}

	copier.CopyWithOption(&customPolicy.Spec.Process, &policy.Process, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	copier.CopyWithOption(&customPolicy.Spec.File, &policy.File, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	copier.CopyWithOption(&customPolicy.Spec.Capabilities, &policy.Network, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	copier.CopyWithOption(&customPolicy.Spec.Network, &policy.Network, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	customPolicy.Spec.Action = policy.Action

	return &customPolicy
}

func GenerateCustomPolicyProfile(policy varmor.Policy, profileName string) string {
	customPolicies := []kubearmortype.SecurityPolicy{}

	initCustomPolicy(&policy.CustomPolicy)
	p := newKubeArmorSecurityPolicy(policy.CustomPolicy)
	customPolicies = append(customPolicies, *p)

	_, rules := kubearmorenforcer.GenerateProfileBody(customPolicies)
	c := []byte(fmt.Sprintf(customPolicyRulesTemplate, profileName, rules))

	return base64.StdEncoding.EncodeToString(c)
}

func GenerateBehaviorModelingProfile(profileName string) string {
	c := []byte(fmt.Sprintf(behaviorModelingTemplate, profileName))
	return base64.StdEncoding.EncodeToString(c)
}
