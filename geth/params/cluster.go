package params

type cluster struct {
	NetworkID   int      `json:"networkID"`
	Discovery   bool     `json:"discovery"`
	StaticNodes []string `json:"staticnodes"`
	BootNodes   []string `json:"bootnodes"`
}

var ropstenCluster = cluster{
	NetworkID: 3,
	StaticNodes: []string{
		"enode://dffef3874011709b12d1e540d83ddb19a9db8614ad9151d05bcf813585e45cbebba5aaea223fe315786c401d8cecb1ad2de9f179680c536ea30311fb21fa934b@188.166.100.178:30303",
		"enode://03f3661686d30509d621dbe5ee2e3082923f25e94fd41a2dd8dd34bb12a0c4e8fbde52247c6c55e86dc209a8e7c4a5ae56058c65f7b01734d3ab73818b44e2a3@188.166.33.47:30303",
	},
}

var rinkebyCluster = cluster{
	NetworkID: 4,
	Discovery: true,
	BootNodes: []string{
		"enode://b29100c8468e3e6604817174a15e4d71627458b0dcdbeea169ab2eb4ab2bbc6f24adbb175826726cec69db8fdba6c0dd60b3da598e530ede562180d300728659@206.189.6.48:30404",
		"enode://1b843c7697f6fc42a1f606fb3cfaac54e025f06789dc20ad9278be3388967cf21e3a1b1e4be51faecd66c2c3adef12e942b4fcdeb8727657abe60636efb6224f@206.189.6.46:30404",
	},
}

var mainnetCluster = cluster{
	NetworkID: 1,
	StaticNodes: []string{
		"enode://3aeaff0868b19e03fabe33e6e0fcc821094e1601be44edd6f45e3f0171ed964e13623e49987bddd6c517304d2a45dfe66da51e47b2e11d59c4b30cd6094db43d@163.172.176.22:30303",
		"enode://687343483ca41132a16c9ab67b49e9997a34ec38ddb6dd60bf45f9a0ea4c50362f902553d813af44ab1cdb246fc384d4c74b4437c15cefe3bb0e87b399dbb5bb@163.172.176.22:30403",
		"enode://2a3d6c1c86546831e5bb2684ff0ed6d931bdacf3c6cd344706452a1e78c41442d38c62317096175dcea6517959f40ac789f76356348e0a17ee53563cbdf2db48@163.172.176.22:30503",
		"enode://71bb01b58165e3262aea2d3b06dbf9abb8d5512d96e5000e7e41ab2138b47be685935d3eb119fc25e1413db00d8db231fd9d59555a1cd75229821559b6a4eb51@51.15.85.243:30303",
		"enode://7afd119c549a7ab02b3f7bd77ef3490b6d660d5c49d0734a0c8bb23195ced4ace0bf5cde673cd5cfd07dd8d759277f3d8408eb73dc3c217bbe00f0027d06eee9@51.15.85.243:30403",
		"enode://da8af0869e4e8047f21c1ac016b94a7b7d8e935dddd28d4272f88a1ceaee7c15e7deec9b6fd195ed3bc43748893111ebf2b2479ff44a8025ab8d598f3c97b589@51.15.85.243:30503",
		"enode://7ebaa6a8ce2547f10e34fab9cc5626b86d67934a86e1fb36145c0b89fcc7b9315dd6d0a8cc5808d11a55bdc14c78ff675ca956dfec53837b4f1a97392b15ec23@51.15.35.110:30303",
	},
}

var defaultClusters = []cluster{ropstenCluster, rinkebyCluster, mainnetCluster}
