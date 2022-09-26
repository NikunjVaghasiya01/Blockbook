//go:build unittest

package koto

import (
	"bytes"
	"encoding/hex"
	"math/big"
	"os"
	"reflect"
	"testing"

	"github.com/martinboehm/btcutil/chaincfg"
	"github.com/NikunjVaghasiya01/blockbook/bchain"
	"github.com/NikunjVaghasiya01/blockbook/bchain/coins/btc"
)

var (
	testTx1, testTx2 bchain.Tx

	testTxPacked1 = "0a2097f944e3558cc784f4013b3753ce9570fe4707893eda724b12eb4c69686113a612970f020000000001036b2048020000001976a9142df466d79cf4be0f7d1091512f1c297e4988fdd188ac000000000100000000000000001392204802000000a8a3d9a54a3fe3b5ae208fa2d96faea96dac6cb76b03bb2e32c5dd892a5d6f6490f05d8f6fb4401228df1be9f22c5ad69706c461bbc9253ffbc3531770e5075e149ec6af7f2cda0cb87862620792340bc425095115adbf0f16f4d3ece3f5c467cbcb02b01ced7192ab644f96fa01f04a7a450a42da2bfad1748634f7bc141467d3c94961ed6250dd23fca62b30b4a2ca6fbad0724372a429bcb97a28954e1681c0bb974877dac26eb2b994eaa23d56ecfdedd93f8331f9432f12adc37da9f049585179b7bbb370b76c6c37a438a20bc3b410a6a72ff8d11408a337a37bbc73dd15df8a34f2c878dc6d6db4f0cee504680fe53e0a158f1e1c82b84e065a4764fc57f8bf75d28899126917a05bf3230036f2d6b38f8a51214d1c2d0588a95e82f0032c2dfa6916c689f8daa648c01517bbf0826d2d4082067b0d17071920eb6dab6c0307603b825f3aae347db349628d4fcaa97a155ef1a2c601170fe825609efa964f0a06700afc135542ea7b06fc989424d7e100652c0ad5be4ca01c1fd676530e6f60606c5feb7de5da0d69544d8b7be8de06b27ba96a1bfa6bc07cc0269982acb722032a938ab36089c0eeda7a4076cd258a1752486a3d52af16db8dff072bcc61f17503185b5bc8aa0ea3a181663bb0ddd3cca1a19293764b01569b9878a60c0ce82e21020751a4ecc1a2a9b9a123042ee8d8a5d4d3f4764b1cdb13d57d2a77c3b56bd89102302d118ebc14969ade27bf83f0e707a97b26c7292a6c20e850abce5fad0ab59d032fec0d6278bbe0f2fb3fbc61697ba10b6ec3f2c4196e46e98dfb65bb28ec6afff81de5a4be7be8c4f56ab4c03043a3cf9987b630ac4b6d8aa74ce8ed5b61040262239172c6450f9ab642e2f2d258c200c3bb4ad69011ef2dfe1dd63d758c86270ac4925b248b9bb4b6c9ff7bf2e56260cba02b2429648dc20eb034c8f9b18e1f6a38b8651c236554546585b4dd0f07fd5ee1696bf792527ce84b22012439300797103f22d3969df725e4414d899ec32a2ebbb857cc911e374e84738f4e007ee5260ff666a286e45c465525e2e3fc5e5e0e9ad82e53d364e4fd355619711c616d508470b997af44f62f283871cd892552128135aadb40c6f8cf69ee72acf349e9f4d33e8673450b9f69d4022a8d886b0cdbaba0798e0bf57b42dedfafb0bbd5495ca1c0030bbf460b48f9a138f6ab748df9046d9995f895062583ce0818e40afb9704653e11d58ca42bd3f60f4e908589ad9144c76067dc433cad13a5bbd9c168691b8c6cddb19d812f3e3f98e2cdd20dbb170936fd5cd2ef0bca72af8931a1b01d6081ffbef5be4416e696a7c762a375b368f71dc31362a4005750992a48e55311dde8d2013180d62e507ffc3e468c4a27acc763a9651b19f37e1ffca7e656225617368e79c1d18f9b14d770993d3d1dc42dcfd9adfa02a8ddf0ebc8fbb850fab307fd1d239cf6ad4e5ff40992dd974bc43fa351ce807cc0036c2f7d80bbd052f496216304fbc63e8d728bf129acaedb0073aff077e584ce04bc1ccf9c91f41f3c8804dd65da4ecdfdba32590e04b4d1b6895dea8edacd1f40313e8a1d4900d0dba54056eee72e3d155e9c67e7a51df581c33cbd39f16549d590ae5387fe2c5ad3484ad5c7da320066b79083c49879e45938b3bbb063726008a2ebb8847c9e57be6ec489c7aaa80f5e8e430040cb8d60298363df850cb7b4e98e97192882d10d2fd96cc490dd18b263d96aac6aa4f5583770e0917fa9b566dd0e0b218c6684007ec10cf11747e8f039fac5250170de2835ab88fea356b6a7d0f5e81ffe9b78d191a745e0237a256a2a840880689d83503b72462e3955b61e22afde947c1f1527ea94151c5b7d3a72ce68979603911c08bcec01097899fd30347be7f2e246f70d2af6a1e29b54988978a91f79b2ed8be76ffd62f79de5418933ed166be919d9bbb7524347d87d31afaed05e71a82b09c18c196b3ba6e226939b375903f7d889422863567203814484af89fd223ce1c959b1fdffaf26461630c630d2bbf99228a096ea6cb0d61df70d24414c76bf9371c4abff0ad257098189af6ea32200fbe092d875aa4d3f72a7ec138439e4b08fb1dcde6a90f25fde1498773e693c9b21c40505d42edcbcaed8a2dc4642750e9df73e169f9986ddb3a57991ac2cb3b540d788e2c2c22c2c51b2d74a98ed59a8cec89ba54342fb9660449a116f8691da60cb447afe4d5e80f37b4669e6007c1cadc41933fb27bab41afd312c37e5cb43715cb4013efcd91221ed06249540b733c05e81131aba75ba0f427d9bd975554b2d49a8048f0b3a84477e75290235fd3bbcaee6c4438ba72299dd960f3f6ee9241f7e399684e894d7bb1c302ecbe24d0f19dca982a82ee44f36211d23b0ea623c9c9f4f527f4e452fd06ebb943cadeea3d7fa42cabd25324bc5851e40f9952823f56b50b97729e6561f2100c2b5922860c6cf447a668324ca931f2f35a5edb7d306f8b8802f98cf67140a3fe73099ab86bb65c439a8593e64816bcd46aa4c254918f4a3a0f3f47b4ebcfb2824703f9a7d163824484ce6fe1852c4ba131ee2635de14822a8cb3782697fecc6f69514edd3f42fcf2751075b838bf14ae91e9dcff517bf3cec4db1b986b4c966a4fa40d38f2ebb7fc60218c397a2d705200028d09a0c3a490a050248206b0310001a1976a9142df466d79cf4be0f7d1091512f1c297e4988fdd188ac22236b31323257767a46415565444b3667353238563376726547673870614a5137624248364000"
	testTxPacked2 = "0a203aebcf5a223450bca3c0312d3d87b6070447e795d09a266a3a01c70e44c7cc4812e1010100000001cbc2c0b14b26f563ceee8201971b2caae2a4f964d0fd91267290c51a6a171411010000006a473044022032dd5d573c3a7f729da1cb9d9ba02a08e05d50b4f74d5aeb7cb22284526f70340220661ca4a192d02684f0b6b52768b9e9ae5fad41b962aa918537b91bba275e92e70121024e98e62782ba44e5677b52b1e4e973a027c7d873915a6d62ba967b2c07467224ffffffff02c0c62d00000000001976a914dd985697513887236c484acc605ece839e2204ac88ac989e8ce0000000001976a91482bfe75940a6d46238f55e258fcae5bef4e847ea88ac0000000018ff98a2d705200028d49a0c3298010a0012201114176a1ac590722691fdd064f9a4e2aa2c1b970182eece63f5264bb1c0c2cb1801226a473044022032dd5d573c3a7f729da1cb9d9ba02a08e05d50b4f74d5aeb7cb22284526f70340220661ca4a192d02684f0b6b52768b9e9ae5fad41b962aa918537b91bba275e92e70121024e98e62782ba44e5677b52b1e4e973a027c7d873915a6d62ba967b2c0746722428ffffffff0f3a470a032dc6c010001a1976a914dd985697513887236c484acc605ece839e2204ac88ac22236b314a334461347236356653616b6571555953616a6f506f74656376633768384861513a480a04e08c9e9810011a1976a91482bfe75940a6d46238f55e258fcae5bef4e847ea88ac22236b31396b7355666462355139584b556a3565645570314451686e6343503868396845374000"
)

func init() {
	testTx1 = bchain.Tx{
		Hex:       "020000000001036b2048020000001976a9142df466d79cf4be0f7d1091512f1c297e4988fdd188ac000000000100000000000000001392204802000000a8a3d9a54a3fe3b5ae208fa2d96faea96dac6cb76b03bb2e32c5dd892a5d6f6490f05d8f6fb4401228df1be9f22c5ad69706c461bbc9253ffbc3531770e5075e149ec6af7f2cda0cb87862620792340bc425095115adbf0f16f4d3ece3f5c467cbcb02b01ced7192ab644f96fa01f04a7a450a42da2bfad1748634f7bc141467d3c94961ed6250dd23fca62b30b4a2ca6fbad0724372a429bcb97a28954e1681c0bb974877dac26eb2b994eaa23d56ecfdedd93f8331f9432f12adc37da9f049585179b7bbb370b76c6c37a438a20bc3b410a6a72ff8d11408a337a37bbc73dd15df8a34f2c878dc6d6db4f0cee504680fe53e0a158f1e1c82b84e065a4764fc57f8bf75d28899126917a05bf3230036f2d6b38f8a51214d1c2d0588a95e82f0032c2dfa6916c689f8daa648c01517bbf0826d2d4082067b0d17071920eb6dab6c0307603b825f3aae347db349628d4fcaa97a155ef1a2c601170fe825609efa964f0a06700afc135542ea7b06fc989424d7e100652c0ad5be4ca01c1fd676530e6f60606c5feb7de5da0d69544d8b7be8de06b27ba96a1bfa6bc07cc0269982acb722032a938ab36089c0eeda7a4076cd258a1752486a3d52af16db8dff072bcc61f17503185b5bc8aa0ea3a181663bb0ddd3cca1a19293764b01569b9878a60c0ce82e21020751a4ecc1a2a9b9a123042ee8d8a5d4d3f4764b1cdb13d57d2a77c3b56bd89102302d118ebc14969ade27bf83f0e707a97b26c7292a6c20e850abce5fad0ab59d032fec0d6278bbe0f2fb3fbc61697ba10b6ec3f2c4196e46e98dfb65bb28ec6afff81de5a4be7be8c4f56ab4c03043a3cf9987b630ac4b6d8aa74ce8ed5b61040262239172c6450f9ab642e2f2d258c200c3bb4ad69011ef2dfe1dd63d758c86270ac4925b248b9bb4b6c9ff7bf2e56260cba02b2429648dc20eb034c8f9b18e1f6a38b8651c236554546585b4dd0f07fd5ee1696bf792527ce84b22012439300797103f22d3969df725e4414d899ec32a2ebbb857cc911e374e84738f4e007ee5260ff666a286e45c465525e2e3fc5e5e0e9ad82e53d364e4fd355619711c616d508470b997af44f62f283871cd892552128135aadb40c6f8cf69ee72acf349e9f4d33e8673450b9f69d4022a8d886b0cdbaba0798e0bf57b42dedfafb0bbd5495ca1c0030bbf460b48f9a138f6ab748df9046d9995f895062583ce0818e40afb9704653e11d58ca42bd3f60f4e908589ad9144c76067dc433cad13a5bbd9c168691b8c6cddb19d812f3e3f98e2cdd20dbb170936fd5cd2ef0bca72af8931a1b01d6081ffbef5be4416e696a7c762a375b368f71dc31362a4005750992a48e55311dde8d2013180d62e507ffc3e468c4a27acc763a9651b19f37e1ffca7e656225617368e79c1d18f9b14d770993d3d1dc42dcfd9adfa02a8ddf0ebc8fbb850fab307fd1d239cf6ad4e5ff40992dd974bc43fa351ce807cc0036c2f7d80bbd052f496216304fbc63e8d728bf129acaedb0073aff077e584ce04bc1ccf9c91f41f3c8804dd65da4ecdfdba32590e04b4d1b6895dea8edacd1f40313e8a1d4900d0dba54056eee72e3d155e9c67e7a51df581c33cbd39f16549d590ae5387fe2c5ad3484ad5c7da320066b79083c49879e45938b3bbb063726008a2ebb8847c9e57be6ec489c7aaa80f5e8e430040cb8d60298363df850cb7b4e98e97192882d10d2fd96cc490dd18b263d96aac6aa4f5583770e0917fa9b566dd0e0b218c6684007ec10cf11747e8f039fac5250170de2835ab88fea356b6a7d0f5e81ffe9b78d191a745e0237a256a2a840880689d83503b72462e3955b61e22afde947c1f1527ea94151c5b7d3a72ce68979603911c08bcec01097899fd30347be7f2e246f70d2af6a1e29b54988978a91f79b2ed8be76ffd62f79de5418933ed166be919d9bbb7524347d87d31afaed05e71a82b09c18c196b3ba6e226939b375903f7d889422863567203814484af89fd223ce1c959b1fdffaf26461630c630d2bbf99228a096ea6cb0d61df70d24414c76bf9371c4abff0ad257098189af6ea32200fbe092d875aa4d3f72a7ec138439e4b08fb1dcde6a90f25fde1498773e693c9b21c40505d42edcbcaed8a2dc4642750e9df73e169f9986ddb3a57991ac2cb3b540d788e2c2c22c2c51b2d74a98ed59a8cec89ba54342fb9660449a116f8691da60cb447afe4d5e80f37b4669e6007c1cadc41933fb27bab41afd312c37e5cb43715cb4013efcd91221ed06249540b733c05e81131aba75ba0f427d9bd975554b2d49a8048f0b3a84477e75290235fd3bbcaee6c4438ba72299dd960f3f6ee9241f7e399684e894d7bb1c302ecbe24d0f19dca982a82ee44f36211d23b0ea623c9c9f4f527f4e452fd06ebb943cadeea3d7fa42cabd25324bc5851e40f9952823f56b50b97729e6561f2100c2b5922860c6cf447a668324ca931f2f35a5edb7d306f8b8802f98cf67140a3fe73099ab86bb65c439a8593e64816bcd46aa4c254918f4a3a0f3f47b4ebcfb2824703f9a7d163824484ce6fe1852c4ba131ee2635de14822a8cb3782697fecc6f69514edd3f42fcf2751075b838bf14ae91e9dcff517bf3cec4db1b986b4c966a4fa40d38f2ebb7fc602",
		Blocktime: 1525189571,
		Time:      1525189571,
		Txid:      "97f944e3558cc784f4013b3753ce9570fe4707893eda724b12eb4c69686113a6",
		LockTime:  0,
		Vin:       []bchain.Vin{},
		Vout: []bchain.Vout{
			{
				ValueSat: *big.NewInt(9800018691),
				N:        0,
				ScriptPubKey: bchain.ScriptPubKey{
					Hex: "76a9142df466d79cf4be0f7d1091512f1c297e4988fdd188ac",
					Addresses: []string{
						"k122WvzFAUeDK6g528V3vreGg8paJQ7bBH6",
					},
				},
			},
		},
	}

	testTx2 = bchain.Tx{
		Hex:       "0100000001cbc2c0b14b26f563ceee8201971b2caae2a4f964d0fd91267290c51a6a171411010000006a473044022032dd5d573c3a7f729da1cb9d9ba02a08e05d50b4f74d5aeb7cb22284526f70340220661ca4a192d02684f0b6b52768b9e9ae5fad41b962aa918537b91bba275e92e70121024e98e62782ba44e5677b52b1e4e973a027c7d873915a6d62ba967b2c07467224ffffffff02c0c62d00000000001976a914dd985697513887236c484acc605ece839e2204ac88ac989e8ce0000000001976a91482bfe75940a6d46238f55e258fcae5bef4e847ea88ac00000000",
		Blocktime: 1525189759,
		Time:      1525189759,
		Txid:      "3aebcf5a223450bca3c0312d3d87b6070447e795d09a266a3a01c70e44c7cc48",
		LockTime:  0,
		Vin: []bchain.Vin{
			{
				ScriptSig: bchain.ScriptSig{
					Hex: "473044022032dd5d573c3a7f729da1cb9d9ba02a08e05d50b4f74d5aeb7cb22284526f70340220661ca4a192d02684f0b6b52768b9e9ae5fad41b962aa918537b91bba275e92e70121024e98e62782ba44e5677b52b1e4e973a027c7d873915a6d62ba967b2c07467224",
				},
				Txid:     "1114176a1ac590722691fdd064f9a4e2aa2c1b970182eece63f5264bb1c0c2cb",
				Vout:     1,
				Sequence: 4294967295,
			},
		},
		Vout: []bchain.Vout{
			{
				ValueSat: *big.NewInt(3000000),
				N:        0,
				ScriptPubKey: bchain.ScriptPubKey{
					Hex: "76a914dd985697513887236c484acc605ece839e2204ac88ac",
					Addresses: []string{
						"k1J3Da4r65fSakeqUYSajoPotecvc7h8HaQ",
					},
				},
			},
			{
				ValueSat: *big.NewInt(3767312024),
				N:        1,
				ScriptPubKey: bchain.ScriptPubKey{
					Hex: "76a91482bfe75940a6d46238f55e258fcae5bef4e847ea88ac",
					Addresses: []string{
						"k19ksUfdb5Q9XKUj5edUp1DQhncCP8h9hE7",
					},
				},
			},
		},
	}
}

func TestMain(m *testing.M) {
	c := m.Run()
	chaincfg.ResetParams()
	os.Exit(c)
}

func TestGetAddrDesc(t *testing.T) {
	type args struct {
		tx     bchain.Tx
		parser *KotoParser
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "koto-1",
			args: args{
				tx:     testTx1,
				parser: NewKotoParser(GetChainParams("main"), &btc.Configuration{}),
			},
		},
		{
			name: "koto-2",
			args: args{
				tx:     testTx2,
				parser: NewKotoParser(GetChainParams("main"), &btc.Configuration{}),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for n, vout := range tt.args.tx.Vout {
				got1, err := tt.args.parser.GetAddrDescFromVout(&vout)
				if err != nil {
					t.Errorf("getAddrDescFromVout() error = %v, vout = %d", err, n)
					return
				}
				got2, err := tt.args.parser.GetAddrDescFromAddress(vout.ScriptPubKey.Addresses[0])
				if err != nil {
					t.Errorf("getAddrDescFromAddress() error = %v, vout = %d", err, n)
					return
				}
				if !bytes.Equal(got1, got2) {
					t.Errorf("Address descriptors mismatch: got1 = %v, got2 = %v", got1, got2)
				}
			}
		})
	}
}

func TestPackTx(t *testing.T) {
	type args struct {
		tx        bchain.Tx
		height    uint32
		blockTime int64
		parser    *KotoParser
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "koto-1",
			args: args{
				tx:        testTx1,
				height:    200016,
				blockTime: 1525189571,
				parser:    NewKotoParser(GetChainParams("main"), &btc.Configuration{}),
			},
			want:    testTxPacked1,
			wantErr: false,
		},
		{
			name: "koto-2",
			args: args{
				tx:        testTx2,
				height:    200020,
				blockTime: 1525189759,
				parser:    NewKotoParser(GetChainParams("main"), &btc.Configuration{}),
			},
			want:    testTxPacked2,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.args.parser.PackTx(&tt.args.tx, tt.args.height, tt.args.blockTime)
			if (err != nil) != tt.wantErr {
				t.Errorf("packTx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			h := hex.EncodeToString(got)
			if !reflect.DeepEqual(h, tt.want) {
				t.Errorf("packTx() = %v, want %v", h, tt.want)
			}
		})
	}
}

func TestUnpackTx(t *testing.T) {
	type args struct {
		packedTx string
		parser   *KotoParser
	}
	tests := []struct {
		name    string
		args    args
		want    *bchain.Tx
		want1   uint32
		wantErr bool
	}{
		{
			name: "koto-1",
			args: args{
				packedTx: testTxPacked1,
				parser:   NewKotoParser(GetChainParams("main"), &btc.Configuration{}),
			},
			want:    &testTx1,
			want1:   200016,
			wantErr: false,
		},
		{
			name: "koto-2",
			args: args{
				packedTx: testTxPacked2,
				parser:   NewKotoParser(GetChainParams("main"), &btc.Configuration{}),
			},
			want:    &testTx2,
			want1:   200020,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, _ := hex.DecodeString(tt.args.packedTx)
			got, got1, err := tt.args.parser.UnpackTx(b)
			if (err != nil) != tt.wantErr {
				t.Errorf("unpackTx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("unpackTx() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("unpackTx() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
