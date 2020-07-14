// Code generated by goa v3.1.3, DO NOT EDIT.
//
// fruitshop HTTP client CLI support package
//
// Command:
// $ goa gen fruitshop/design

package cli

import (
	"flag"
	"fmt"
	cartc "fruitshop/gen/http/cart/client"
	discountc "fruitshop/gen/http/discount/client"
	fruitc "fruitshop/gen/http/fruit/client"
	paymentc "fruitshop/gen/http/payment/client"
	userc "fruitshop/gen/http/user/client"
	"net/http"
	"os"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `user (add|get|show)
fruit (get|show)
cart (add|remove|get)
payment (add|get)
discount get
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` user add --body '{
      "ID": "Ab enim."
   }' --user-id "Incidunt quae quia officia rerum est."` + "\n" +
		os.Args[0] + ` fruit get --body '{
      "cost": 0.9509512832245888
   }' --name "Doloremque corrupti quo et."` + "\n" +
		os.Args[0] + ` cart add --body '{
      "costPerItem": 0.11945221888838935,
      "count": 8003916626658546661,
      "name": "Officiis officiis et cumque eum est consequatur.",
      "totalCost": 0.2826886157163321
   }' --user-id "Eos blanditiis ut nesciunt."` + "\n" +
		os.Args[0] + ` payment add --body '{
      "ID": "Id earum explicabo tenetur sit.",
      "amount": 0.6465372050495056
   }' --user-id "Quidem velit quidem."` + "\n" +
		os.Args[0] + ` discount get --user-id "Sit nesciunt cumque et provident eaque est."` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, interface{}, error) {
	var (
		userFlags = flag.NewFlagSet("user", flag.ContinueOnError)

		userAddFlags      = flag.NewFlagSet("add", flag.ExitOnError)
		userAddBodyFlag   = userAddFlags.String("body", "REQUIRED", "")
		userAddUserIDFlag = userAddFlags.String("user-id", "REQUIRED", "userId")

		userGetFlags      = flag.NewFlagSet("get", flag.ExitOnError)
		userGetUserIDFlag = userGetFlags.String("user-id", "REQUIRED", "userId")

		userShowFlags = flag.NewFlagSet("show", flag.ExitOnError)

		fruitFlags = flag.NewFlagSet("fruit", flag.ContinueOnError)

		fruitGetFlags    = flag.NewFlagSet("get", flag.ExitOnError)
		fruitGetBodyFlag = fruitGetFlags.String("body", "REQUIRED", "")
		fruitGetNameFlag = fruitGetFlags.String("name", "REQUIRED", "name")

		fruitShowFlags = flag.NewFlagSet("show", flag.ExitOnError)

		cartFlags = flag.NewFlagSet("cart", flag.ContinueOnError)

		cartAddFlags      = flag.NewFlagSet("add", flag.ExitOnError)
		cartAddBodyFlag   = cartAddFlags.String("body", "REQUIRED", "")
		cartAddUserIDFlag = cartAddFlags.String("user-id", "REQUIRED", "ID of the user")

		cartRemoveFlags      = flag.NewFlagSet("remove", flag.ExitOnError)
		cartRemoveBodyFlag   = cartRemoveFlags.String("body", "REQUIRED", "")
		cartRemoveUserIDFlag = cartRemoveFlags.String("user-id", "REQUIRED", "ID of the user")

		cartGetFlags      = flag.NewFlagSet("get", flag.ExitOnError)
		cartGetUserIDFlag = cartGetFlags.String("user-id", "REQUIRED", "ID")

		paymentFlags = flag.NewFlagSet("payment", flag.ContinueOnError)

		paymentAddFlags      = flag.NewFlagSet("add", flag.ExitOnError)
		paymentAddBodyFlag   = paymentAddFlags.String("body", "REQUIRED", "")
		paymentAddUserIDFlag = paymentAddFlags.String("user-id", "REQUIRED", "userId of the user")

		paymentGetFlags      = flag.NewFlagSet("get", flag.ExitOnError)
		paymentGetBodyFlag   = paymentGetFlags.String("body", "REQUIRED", "")
		paymentGetUserIDFlag = paymentGetFlags.String("user-id", "REQUIRED", "userId")

		discountFlags = flag.NewFlagSet("discount", flag.ContinueOnError)

		discountGetFlags      = flag.NewFlagSet("get", flag.ExitOnError)
		discountGetUserIDFlag = discountGetFlags.String("user-id", "REQUIRED", "userId")
	)
	userFlags.Usage = userUsage
	userAddFlags.Usage = userAddUsage
	userGetFlags.Usage = userGetUsage
	userShowFlags.Usage = userShowUsage

	fruitFlags.Usage = fruitUsage
	fruitGetFlags.Usage = fruitGetUsage
	fruitShowFlags.Usage = fruitShowUsage

	cartFlags.Usage = cartUsage
	cartAddFlags.Usage = cartAddUsage
	cartRemoveFlags.Usage = cartRemoveUsage
	cartGetFlags.Usage = cartGetUsage

	paymentFlags.Usage = paymentUsage
	paymentAddFlags.Usage = paymentAddUsage
	paymentGetFlags.Usage = paymentGetUsage

	discountFlags.Usage = discountUsage
	discountGetFlags.Usage = discountGetUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "user":
			svcf = userFlags
		case "fruit":
			svcf = fruitFlags
		case "cart":
			svcf = cartFlags
		case "payment":
			svcf = paymentFlags
		case "discount":
			svcf = discountFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "user":
			switch epn {
			case "add":
				epf = userAddFlags

			case "get":
				epf = userGetFlags

			case "show":
				epf = userShowFlags

			}

		case "fruit":
			switch epn {
			case "get":
				epf = fruitGetFlags

			case "show":
				epf = fruitShowFlags

			}

		case "cart":
			switch epn {
			case "add":
				epf = cartAddFlags

			case "remove":
				epf = cartRemoveFlags

			case "get":
				epf = cartGetFlags

			}

		case "payment":
			switch epn {
			case "add":
				epf = paymentAddFlags

			case "get":
				epf = paymentGetFlags

			}

		case "discount":
			switch epn {
			case "get":
				epf = discountGetFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "user":
			c := userc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "add":
				endpoint = c.Add()
				data, err = userc.BuildAddPayload(*userAddBodyFlag, *userAddUserIDFlag)
			case "get":
				endpoint = c.Get()
				data, err = userc.BuildGetPayload(*userGetUserIDFlag)
			case "show":
				endpoint = c.Show()
				data = nil
			}
		case "fruit":
			c := fruitc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "get":
				endpoint = c.Get()
				data, err = fruitc.BuildGetPayload(*fruitGetBodyFlag, *fruitGetNameFlag)
			case "show":
				endpoint = c.Show()
				data = nil
			}
		case "cart":
			c := cartc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "add":
				endpoint = c.Add()
				data, err = cartc.BuildAddPayload(*cartAddBodyFlag, *cartAddUserIDFlag)
			case "remove":
				endpoint = c.Remove()
				data, err = cartc.BuildRemovePayload(*cartRemoveBodyFlag, *cartRemoveUserIDFlag)
			case "get":
				endpoint = c.Get()
				data, err = cartc.BuildGetPayload(*cartGetUserIDFlag)
			}
		case "payment":
			c := paymentc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "add":
				endpoint = c.Add()
				data, err = paymentc.BuildAddPayload(*paymentAddBodyFlag, *paymentAddUserIDFlag)
			case "get":
				endpoint = c.Get()
				data, err = paymentc.BuildGetPayload(*paymentGetBodyFlag, *paymentGetUserIDFlag)
			}
		case "discount":
			c := discountc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "get":
				endpoint = c.Get()
				data, err = discountc.BuildGetPayload(*discountGetUserIDFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// userUsage displays the usage of the user command and its subcommands.
func userUsage() {
	fmt.Fprintf(os.Stderr, `The user service allows access to user members
Usage:
    %s [globalflags] user COMMAND [flags]

COMMAND:
    add: Add implements add.
    get: Get implements get.
    show: Show implements show.

Additional help:
    %s user COMMAND --help
`, os.Args[0], os.Args[0])
}
func userAddUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] user add -body JSON -user-id STRING

Add implements add.
    -body JSON: 
    -user-id STRING: userId

Example:
    `+os.Args[0]+` user add --body '{
      "ID": "Ab enim."
   }' --user-id "Incidunt quae quia officia rerum est."
`, os.Args[0])
}

func userGetUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] user get -user-id STRING

Get implements get.
    -user-id STRING: userId

Example:
    `+os.Args[0]+` user get --user-id "Molestiae corporis suscipit quidem."
`, os.Args[0])
}

func userShowUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] user show

Show implements show.

Example:
    `+os.Args[0]+` user show
`, os.Args[0])
}

// fruitUsage displays the usage of the fruit command and its subcommands.
func fruitUsage() {
	fmt.Fprintf(os.Stderr, `The user service allows access to fruits
Usage:
    %s [globalflags] fruit COMMAND [flags]

COMMAND:
    get: Get implements get.
    show: Show implements show.

Additional help:
    %s fruit COMMAND --help
`, os.Args[0], os.Args[0])
}
func fruitGetUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] fruit get -body JSON -name STRING

Get implements get.
    -body JSON: 
    -name STRING: name

Example:
    `+os.Args[0]+` fruit get --body '{
      "cost": 0.9509512832245888
   }' --name "Doloremque corrupti quo et."
`, os.Args[0])
}

func fruitShowUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] fruit show

Show implements show.

Example:
    `+os.Args[0]+` fruit show
`, os.Args[0])
}

// cartUsage displays the usage of the cart command and its subcommands.
func cartUsage() {
	fmt.Fprintf(os.Stderr, `The cart service allows to manage the state of the cart
Usage:
    %s [globalflags] cart COMMAND [flags]

COMMAND:
    add: Add implements add.
    remove: Remove implements remove.
    get: Get implements get.

Additional help:
    %s cart COMMAND --help
`, os.Args[0], os.Args[0])
}
func cartAddUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] cart add -body JSON -user-id STRING

Add implements add.
    -body JSON: 
    -user-id STRING: ID of the user

Example:
    `+os.Args[0]+` cart add --body '{
      "costPerItem": 0.11945221888838935,
      "count": 8003916626658546661,
      "name": "Officiis officiis et cumque eum est consequatur.",
      "totalCost": 0.2826886157163321
   }' --user-id "Eos blanditiis ut nesciunt."
`, os.Args[0])
}

func cartRemoveUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] cart remove -body JSON -user-id STRING

Remove implements remove.
    -body JSON: 
    -user-id STRING: ID of the user

Example:
    `+os.Args[0]+` cart remove --body '{
      "costPerItem": 0.9308342602059104,
      "count": 1819382036291031517,
      "name": "Est officiis ut.",
      "totalCost": 0.2421495069964654
   }' --user-id "Et vel ad labore."
`, os.Args[0])
}

func cartGetUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] cart get -user-id STRING

Get implements get.
    -user-id STRING: ID

Example:
    `+os.Args[0]+` cart get --user-id "Labore ratione numquam."
`, os.Args[0])
}

// paymentUsage displays the usage of the payment command and its subcommands.
func paymentUsage() {
	fmt.Fprintf(os.Stderr, `The cart service allows to manage the state of the cart
Usage:
    %s [globalflags] payment COMMAND [flags]

COMMAND:
    add: Add implements add.
    get: Get implements get.

Additional help:
    %s payment COMMAND --help
`, os.Args[0], os.Args[0])
}
func paymentAddUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] payment add -body JSON -user-id STRING

Add implements add.
    -body JSON: 
    -user-id STRING: userId of the user

Example:
    `+os.Args[0]+` payment add --body '{
      "ID": "Id earum explicabo tenetur sit.",
      "amount": 0.6465372050495056
   }' --user-id "Quidem velit quidem."
`, os.Args[0])
}

func paymentGetUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] payment get -body JSON -user-id STRING

Get implements get.
    -body JSON: 
    -user-id STRING: userId

Example:
    `+os.Args[0]+` payment get --body '{
      "ID": "Quaerat atque."
   }' --user-id "Nostrum sed."
`, os.Args[0])
}

// discountUsage displays the usage of the discount command and its subcommands.
func discountUsage() {
	fmt.Fprintf(os.Stderr, `Discounts applied on the cart
Usage:
    %s [globalflags] discount COMMAND [flags]

COMMAND:
    get: Get implements get.

Additional help:
    %s discount COMMAND --help
`, os.Args[0], os.Args[0])
}
func discountGetUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] discount get -user-id STRING

Get implements get.
    -user-id STRING: userId

Example:
    `+os.Args[0]+` discount get --user-id "Sit nesciunt cumque et provident eaque est."
`, os.Args[0])
}
