package encoding

import (
	"encoding/hex"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"gopkg.in/go-playground/assert.v1"
)

func Test_EncodeABISignalProof(t *testing.T) {
	// nolint: lll
	accProof1, _ := hex.DecodeString("f90211a0f776494ecffe03ad2af2426ee4fb5ae66c012c03ea3955d5b40a52fe7c6b8df6a0ef4a67411e4accae9bbb05bbcc17d036432de4e3ee0b0d02a04ee6a67db9c1a5a06c4c8f2b3206553ec83f1081a6742edff2f7b7830d3ba4166872c4f3f32bae0ba0aa2e8d96490616498fa142ba7c6a628e2839b9b182c9995992ecf1c1d346784fa0f8d7262d0abdf5d084338a640322c4dd45b1484a630331922afa60256917058aa0e49099e467972084cd70025874c6ae0a6841159aec02f39384de2d9ae7460986a02ad2a355fe840ae1bfe7843616d6bb027c1fd478918cc956e0ca7b61ec0044d0a061ec8e714d951772d74765ae997bd25f3982ea8a0c4dda8baa1e47adbd5a3975a039e5f7e8298126da1e7d6f42b1f20846d68ff40dc857b0dd251d2791c0345589a065f4d771a35d862a10b2a2da6d7d6c1e082cd59280cbab1c8e6aaf89069e72dca07851c1d2a801228a1adb24b61a5b4e1d4e12c5043a0af1b1790cd79b281c43a0a031acf4e7bd8cfb52c19cb512fb2578f416cbcf28cd3b4387d7b6107a17d4ff31a09e8def26ad9cf07bfbba2cf4a3fad0fe4075030b9a67acbecbfa5b7bdc8620cca07174a59e7f5b71a1a20f8c4a232100c1303378a7eb5dd523904d5305e20ee7dca0199ae8cd6d2b09717d59b5106ef31b451b30d777296b4afb43db9327117471eca09193465aefe93e209d05f892f77002dd5c2e74d4eed502ebae3acb7db1e33acc80")

	// nolint: lll
	accProof2, _ := hex.DecodeString("f8f1a0b4df476cf7a24306eb6b96d0a8fe974b88e035d93e6716cf56e3f6bacee15c88808080a09265834cbd374b79cc4f97cdab16770f4434bce18ff9724970bff1c91d54f2f1a0e13b8847e8cfc7fd465a61a64b732d940d8f87ed65828e1004e46fb4256aba48a05a7b35dc9c135fc4df95c4b02179719a3505d201fb590212cf41c31eae1ceb6980a07c46a7878ac08b149792973e8a17982311f3c997624aed0fde76ba45d9a1ba41a0ebfe5b284e549d79c976f878d93266d238d9ea3254d0ef9c199d2c5a12067de48080a05d5e60bc5b531718ba65199b06b68087c640472ef4b3afb7e4b857280edbf2e580808080")

	// nolint: lll
	sProof1, _ := hex.DecodeString("f90211a062880ba06fb396ad4e16f01d22ca7a1ae677e3fe428817928db30e2cae96b97ca0aa57a805600be2304ffdd527bd3dc9e91161233dc59afb71c1aab542eafe70caa03bc6c86a4c6b49305b95b3812a162f7e6bec891b287ef291e9c468071ef0c4ada08ac85ec9872d9e6f5b4c6e72154158df88d44163457cf0bbf2134e385c871a4ea0f35f3c83fbd9da738bbfea1bc87b073d3b64abdecb6294b61cf3eb535eabefdea0905c9b0e1755d389f306f026ccb71f8f7e54cd68720cc1c76682504eeb7bceaea06867477d77649657f698380e66844a7ed14818e8aad9f4ac5748963ede640e0aa0caa272deb3227cb8b0765a718ac85bbc7ee93a04bc0a2cb9c5509c9394470eb3a01689508cc26d870b0183c13bee39ed5342bf32464813f872d7ea4e5bc5f79845a0b578886ee673adcdf7b219cd13d9d641f8e15dd9ec6c9345435e7968bc6bcc82a0fbd86d32d6c60089268373be401211c3b606efeb550659b9e43458007dce2eb6a035d73d30ad77c85ef247ab8473f71314a9d175c1e9a0ce73a78a103a3766f54ca0c08386bed5af43c7cadb42d9df7ba3b94886f483e8a2e66aaef7391a15ab51cba002ce1e689b6193a6d3a8c822b0b0076dfdf732fd045a4dc122ec3879fe3de70ea0db27c27a802c40acbde50e5c357e5327a91242e6550fe461eec10ac136ddddcea0ad6d871b4c62042c68f0ecfdb986a24ea7d850563bbd3d27f6916bc3ddd170a480")

	// nolint: lll
	sProof2, _ := hex.DecodeString("f90211a05c9b8f83e3c03e07271225e2ebce1cbe9e7db3b14d2724ec6efe9cf8fce6fc06a0dbd4cd41e027eefe208271111ea3e99cb39b4645e7e166d084d62f427a9313ada0cc65078735257beecceb9c74985901fa16e8e9fb228ce6aaa62aedb282a1795fa012f4c2ae88c8f0396048da6a095d0fa2c8b86398651cd37a72d68d88d25ff19ea037cda349771733bba3681eda450fee72f5e3dcbb6b8f2acf4a2bd145d0bfad6da0ef1359be1a9f658e580c968b92029dbf62ce7a56932c10acce28b25bf7206665a037d9790673a2be78a1555bee7d37ab10d1b8d94d1f12bb011b7cc7257bf13004a0dd9b4774c203afaaeb098ab623ce32f1df6f8ff0ac1bbcb78e358b7a242cd19aa0dde51d1f37baae98d02b2e35c81030f17407fc31304ab72cf999bb2c7e8abff3a0f8672c12a366e074d6f42c2c7b0c5cc010bc4ec703c65e3b58c4fbfee18e89c2a057ba424e40bd1c6a8e7d494703f392e834d8ca7696759e2c0216ebd18bcf662fa01eafd299e8a772c056e6919eeb67bf7e1098129855234e942cfc18aaf364d39ea0df6b60bdf553e1511f445fdcf1fb7aadc23bf390eeb11145c9e2742552c2ed6da02e79f5afb8c177c40737cea4aed39fe3c0269f5a8989e02c07a0135594b83bb1a035535dac85afa0e4848c0186cc8687bc7d2de0215b97ea43e65c8e4da0a52517a08ce682327123eb41b4d49ef283ffe11d1da1b9d7163e892b775a63dd31072ec080")

	s := []HopProof{
		{
			ChainID:      1,
			BlockID:      1,
			RootHash:     crypto.Keccak256Hash([]byte("ROOT_HASH")),
			CacheOption:  0,
			AccountProof: [][]byte{accProof1, accProof2},
			StorageProof: [][]byte{sProof1, sProof2},
		},
	}

	// nolint: lll
	want := "0x00000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000001da52c4b137f7596e26ee4d4d6f70b13c851cdbe59d4365e4b7aa16e8c893ca63000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c000000000000000000000000000000000000000000000000000000000000004800000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000002800000000000000000000000000000000000000000000000000000000000000214f90211a0f776494ecffe03ad2af2426ee4fb5ae66c012c03ea3955d5b40a52fe7c6b8df6a0ef4a67411e4accae9bbb05bbcc17d036432de4e3ee0b0d02a04ee6a67db9c1a5a06c4c8f2b3206553ec83f1081a6742edff2f7b7830d3ba4166872c4f3f32bae0ba0aa2e8d96490616498fa142ba7c6a628e2839b9b182c9995992ecf1c1d346784fa0f8d7262d0abdf5d084338a640322c4dd45b1484a630331922afa60256917058aa0e49099e467972084cd70025874c6ae0a6841159aec02f39384de2d9ae7460986a02ad2a355fe840ae1bfe7843616d6bb027c1fd478918cc956e0ca7b61ec0044d0a061ec8e714d951772d74765ae997bd25f3982ea8a0c4dda8baa1e47adbd5a3975a039e5f7e8298126da1e7d6f42b1f20846d68ff40dc857b0dd251d2791c0345589a065f4d771a35d862a10b2a2da6d7d6c1e082cd59280cbab1c8e6aaf89069e72dca07851c1d2a801228a1adb24b61a5b4e1d4e12c5043a0af1b1790cd79b281c43a0a031acf4e7bd8cfb52c19cb512fb2578f416cbcf28cd3b4387d7b6107a17d4ff31a09e8def26ad9cf07bfbba2cf4a3fad0fe4075030b9a67acbecbfa5b7bdc8620cca07174a59e7f5b71a1a20f8c4a232100c1303378a7eb5dd523904d5305e20ee7dca0199ae8cd6d2b09717d59b5106ef31b451b30d777296b4afb43db9327117471eca09193465aefe93e209d05f892f77002dd5c2e74d4eed502ebae3acb7db1e33acc8000000000000000000000000000000000000000000000000000000000000000000000000000000000000000f3f8f1a0b4df476cf7a24306eb6b96d0a8fe974b88e035d93e6716cf56e3f6bacee15c88808080a09265834cbd374b79cc4f97cdab16770f4434bce18ff9724970bff1c91d54f2f1a0e13b8847e8cfc7fd465a61a64b732d940d8f87ed65828e1004e46fb4256aba48a05a7b35dc9c135fc4df95c4b02179719a3505d201fb590212cf41c31eae1ceb6980a07c46a7878ac08b149792973e8a17982311f3c997624aed0fde76ba45d9a1ba41a0ebfe5b284e549d79c976f878d93266d238d9ea3254d0ef9c199d2c5a12067de48080a05d5e60bc5b531718ba65199b06b68087c640472ef4b3afb7e4b857280edbf2e580808080000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000002800000000000000000000000000000000000000000000000000000000000000214f90211a062880ba06fb396ad4e16f01d22ca7a1ae677e3fe428817928db30e2cae96b97ca0aa57a805600be2304ffdd527bd3dc9e91161233dc59afb71c1aab542eafe70caa03bc6c86a4c6b49305b95b3812a162f7e6bec891b287ef291e9c468071ef0c4ada08ac85ec9872d9e6f5b4c6e72154158df88d44163457cf0bbf2134e385c871a4ea0f35f3c83fbd9da738bbfea1bc87b073d3b64abdecb6294b61cf3eb535eabefdea0905c9b0e1755d389f306f026ccb71f8f7e54cd68720cc1c76682504eeb7bceaea06867477d77649657f698380e66844a7ed14818e8aad9f4ac5748963ede640e0aa0caa272deb3227cb8b0765a718ac85bbc7ee93a04bc0a2cb9c5509c9394470eb3a01689508cc26d870b0183c13bee39ed5342bf32464813f872d7ea4e5bc5f79845a0b578886ee673adcdf7b219cd13d9d641f8e15dd9ec6c9345435e7968bc6bcc82a0fbd86d32d6c60089268373be401211c3b606efeb550659b9e43458007dce2eb6a035d73d30ad77c85ef247ab8473f71314a9d175c1e9a0ce73a78a103a3766f54ca0c08386bed5af43c7cadb42d9df7ba3b94886f483e8a2e66aaef7391a15ab51cba002ce1e689b6193a6d3a8c822b0b0076dfdf732fd045a4dc122ec3879fe3de70ea0db27c27a802c40acbde50e5c357e5327a91242e6550fe461eec10ac136ddddcea0ad6d871b4c62042c68f0ecfdb986a24ea7d850563bbd3d27f6916bc3ddd170a4800000000000000000000000000000000000000000000000000000000000000000000000000000000000000214f90211a05c9b8f83e3c03e07271225e2ebce1cbe9e7db3b14d2724ec6efe9cf8fce6fc06a0dbd4cd41e027eefe208271111ea3e99cb39b4645e7e166d084d62f427a9313ada0cc65078735257beecceb9c74985901fa16e8e9fb228ce6aaa62aedb282a1795fa012f4c2ae88c8f0396048da6a095d0fa2c8b86398651cd37a72d68d88d25ff19ea037cda349771733bba3681eda450fee72f5e3dcbb6b8f2acf4a2bd145d0bfad6da0ef1359be1a9f658e580c968b92029dbf62ce7a56932c10acce28b25bf7206665a037d9790673a2be78a1555bee7d37ab10d1b8d94d1f12bb011b7cc7257bf13004a0dd9b4774c203afaaeb098ab623ce32f1df6f8ff0ac1bbcb78e358b7a242cd19aa0dde51d1f37baae98d02b2e35c81030f17407fc31304ab72cf999bb2c7e8abff3a0f8672c12a366e074d6f42c2c7b0c5cc010bc4ec703c65e3b58c4fbfee18e89c2a057ba424e40bd1c6a8e7d494703f392e834d8ca7696759e2c0216ebd18bcf662fa01eafd299e8a772c056e6919eeb67bf7e1098129855234e942cfc18aaf364d39ea0df6b60bdf553e1511f445fdcf1fb7aadc23bf390eeb11145c9e2742552c2ed6da02e79f5afb8c177c40737cea4aed39fe3c0269f5a8989e02c07a0135594b83bb1a035535dac85afa0e4848c0186cc8687bc7d2de0215b97ea43e65c8e4da0a52517a08ce682327123eb41b4d49ef283ffe11d1da1b9d7163e892b775a63dd31072ec080000000000000000000000000"
	proof, err := EncodeHopProofs(s)
	assert.Equal(t, nil, err)
	assert.Equal(t, hexutil.Encode(proof), want)
}