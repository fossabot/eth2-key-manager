package account_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/bloxapp/KeyVault/cli/cmd"
	"github.com/bloxapp/KeyVault/cli/util/printer"
)

func TestAccountCreate(t *testing.T) {
	t.Run("Successfully create account", func(t *testing.T) {
		var output bytes.Buffer
		cmd.ResultPrinter = printer.New(&output)
		cmd.RootCmd.SetArgs([]string{
			"wallet",
			"account",
			"create",
			"--seed=0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1fff",
			"--storage=7b226163636f756e7473223a2237623232333233393631333133323330333233323264333733323633333532643334333833343635326433393334333436323264363333363636333636363634363633383330333933343333323233613762323236323631373336353431363336333666373536653734353036313734363832323361323232663331323232633232363936343232336132323332333936313331333233303332333232643337333236333335326433343338333436353264333933343334363232643633333636363336363636343636333833303339333433333232326332323665363136643635323233613232363136333633333232323263323237363631366336393634363137343639366636653462363537393232336137623232363936343232336132323338333533393337333633343634363332643634363236363339326433343330333533333264333933353634363232643633363133383334363336333633333436323633333933323232326332323730363137343638323233613232366432663331333233333338333132663333333633303330326633313266333032663330323232633232373037323639373634623635373932323361323233333332333533383633333833323635333033363335333036363339333633343337363333313336333536313332333136343631333833333332363133333631333336363334333336353631363233353334333933343636363633393334333736343634333433323337333836333330363133353332363136313332333836343232376432633232373736393734363836343732363137373631366335303735363234623635373932323361323233393337333633333636363233383633363636343330363336363632333833333335333336353330333736353338333533343330333033363636363533383632363433333635363436333338333236363335363536353333333736313633333133363330333236323635333633343330333733393632333336313632363536363631333133313335363433373335333936333635333333323636363536333635333633323631333936323338333933323635333433303632333033353331333232323764326332323333333733323332333436333331333432643339333836343334326433343333363333343264363233303337363532643333333233333330333733323636333233373339333736363232336137623232363236313733363534313633363336663735366537343530363137343638323233613232326633303232326332323639363432323361323233333337333233323334363333313334326433393338363433343264333433333633333432643632333033373635326433333332333333303337333236363332333733393337363632323263323236653631366436353232336132323631363336333331323232633232373636313663363936343631373436393666366534623635373932323361376232323639363432323361323233393633363633373332363336353636326433313632333833353264333433393636363232643631363333363332326433373338363533333339333536343338363133393332333732323263323237303631373436383232336132323664326633313332333333383331326633333336333033303266333032663330326633303232326332323730373236393736346236353739323233613232333233333636363433343336333436333331333233323634333736363631333836333339363333383635333433363634333733313330363136353334333733383631363233393332333036333338363333303335333833373635333833363335333533363631363133393336333833313339333136343335333233313330363532323764326332323737363937343638363437323631373736313663353037353632346236353739323233613232363233333332333333353333333736323332333833363337363433393636333236323631363533303336333836363339333336353337333536313339363533323635333136333338363433353339333436353333333633393336363333333334363436333338333033313330363436333334333033333635363136353635363136363334333333373335333636313334333433303636363333383332363533313633333636363334333536333336363533383333333433383333333433333636323237643764222c226174744d656d6f7279223a2237623764222c2270726f706f73616c4d656d6f7279223a2237623764222c2277616c6c6574223a22376232323639363432323361323236343337333633303635363436313631326433373335363333343264333436343339333332643338363133333635326436343339333433313634333236353332333133383635333232323263323236393665363436353738346436313730373036353732323233613762323236313633363333313232336132323333333733323332333436333331333432643339333836343334326433343333363333343264363233303337363532643333333233333330333733323636333233373339333736363232326332323631363336333332323233613232333233393631333133323330333233323264333733323633333532643334333833343635326433393334333436323264363333363636333636363634363633383330333933343333323237643263323237343739373036353232336132323438343432323764227d",
		})
		err := cmd.RootCmd.Execute()
		require.NoError(t, err)
	})

	t.Run("Fail to HEX decode seed", func(t *testing.T) {
		var output bytes.Buffer
		cmd.ResultPrinter = printer.New(&output)
		cmd.RootCmd.SetArgs([]string{
			"wallet",
			"account",
			"create",
			"--seed=01213",
			"--storage=01213",
		})
		err := cmd.RootCmd.Execute()
		require.Error(t, err)
		require.EqualError(t, err, "failed to HEX decode seed: encoding/hex: odd length hex string")
	})

	t.Run("Fail to JSON un-marshal", func(t *testing.T) {
		var output bytes.Buffer
		cmd.ResultPrinter = printer.New(&output)
		cmd.RootCmd.SetArgs([]string{
			"wallet",
			"account",
			"create",
			"--seed=b5b0177798165f506de1d46e8e5dd131c708c109800c0e0ce7199aec6572f405",
			"--storage=7b226163636f756e7473223a2237623764222c226174744d656d6f7279223a2237623764222c2270726f706f",
		})
		err := cmd.RootCmd.Execute()
		require.Error(t, err)
		require.EqualError(t, err, "failed to JSON un-marshal storage: unexpected end of JSON input")
	})
}
