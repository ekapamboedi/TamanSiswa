package config

type MenuAccessBody struct {
	Url      string
	MainMenu string
	SubMenu  string
}

var MenuAccesses = []MenuAccessBody{
	{
		Url:      "files",
		MainMenu: "Files",
		SubMenu:  "Files",
	},
	{
		Url:      "payrolls",
		MainMenu: "Payroll",
		SubMenu:  "Payroll",
	},
	{
		Url:      "salaries",
		MainMenu: "Company",
		SubMenu:  "Employee",
	},
}

const (
	PayrollDeductions = `<th style="width: 150px; font-weight: 100; text-align: left; font-size: 12px">%s</th>
            <td style="width: 150px; font-size: 12px">: Rp %s</td>`
	PayrollProperties = `<th style="width: 150px; font-weight: 100; text-align: left; font-size: 12px">%s</th>
            <td style="width: 150px; font-size: 12px">: Rp %s</td>`
	PayrollHtml = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Payroll New Table</title>
    <style type="text/css">
        * {
          margin: 0 auto;
          padding: 0;
          /* padding: 10px 50px; */
          text-indent: 0;
          box-sizing: border-box;
          /* min-height: 100vh; */
        }
    </style>
</head>
<body>
    <table style="width: 600px; margin-top: 100px;">
        <tr>
            <td colspan="4" style="text-align: left;">
                <img style="width: fit-content; height: 70px;" src="%s" />
            </td>
        </tr>
        <tr style="height: 50px;">
            <td colspan="4" style="text-align: left; font-weight: 900; font-size: 32px;">%s</td>
        </tr>
        <tr style="height: 20px" />

        <tr>
            <th style="width: 150px; font-weight: 100; text-align: left;">Employee Name</th>
            <td style="width: 150px;">: %s</td>
            <th style="width: 150px; font-weight: 100; text-align: left">Employee ID</th>
            <td style="width: 150px;">: %s</td>
        </tr>

        <tr style="height: 10px" />

        <tr>
            <th style="width: 150px; font-weight: 100; text-align: left">Department</th>
            <td style="width: 150px;">: %s</td>
            <th style="width: 150px; font-weight: 100; text-align: left">Position</th>
            <td style="width: 150px;">: %s</td>
        </tr>

        <tr style="height: 10px" />

        <tr>
            <th style="width: 150px; font-weight: 100; text-align: left">Bank Name</th>
            <td style="width: 150px;">: %s</td>
            <th style="width: 150px; font-weight: 100; text-align: left">Bank Account</th>
            <td style="width: 150px;">: %s</td>
        </tr>

    </table>

    <table style="border-collapse: collapse; width: 600px; margin-top: 24px; margin-bottom: 8px">
        <tr>
            <td style="border-bottom: 1px solid black; border-top: 1px solid black; text-transform: uppercase; height: 30px; ">income</td>
        </tr>
    </table>

    <table style="width: 600px;">
        %s

        <table style="border-collapse: collapse; width: 600px; margin-top: 16px; margin-bottom: 8px">
            <tr>
                <td style="border-bottom: 1px solid black; border-top: 1px solid black; text-transform: uppercase; height: 30px; font-weight: 900;">total income</td>
                <td style="border-bottom: 1px solid black; border-top: 1px solid black; text-transform: uppercase; height: 30px; font-weight: 900; text-align: right;">Rp %s</td>
            </tr>
        </table>

    </table>

    <table style="border-collapse: collapse; width: 600px; margin-top: 24px; margin-bottom: 8px">
        <tr>
            <td style="border-bottom: 1px solid black; border-top: 1px solid black; text-transform: uppercase; height: 30px; ">deduction</td>
        </tr>
    </table>

    <table style="width: 600px;">
        %s

        <table style="border-collapse: collapse; width: 600px; margin-top: 16px; margin-bottom: 8px">
            <tr>
                <td style="border-bottom: 1px solid black; border-top: 1px solid black; text-transform: uppercase; height: 30px; font-weight: 900;">total deduction</td>
                <td style="border-bottom: 1px solid black; border-top: 1px solid black; text-transform: uppercase; height: 30px; font-weight: 900; text-align: right;">Rp %s</td>
            </tr>
        </table>

    </table>

    <table style="border-collapse: collapse; width: 600px; margin-top: 16px; margin-bottom: 8px; background-color: #6B7280 ; font-size: 20px; color: white;">
        <tr>
            <td style="border-bottom: 1px solid black; border-top: 1px solid black; text-transform: uppercase; height: 30px; font-weight: 900;">take home pay</td>
            <td style="border-bottom: 1px solid black; border-top: 1px solid black; text-transform: uppercase; height: 30px; font-weight: 900; text-align: right;">Rp %s</td>
        </tr>

    </table>

    <table style="width: 600px; margin-bottom: 12px; margin-top: 24px;">
        <tr>
            <td style="text-align: right;">
                %s, %s
            </td>
        </tr>
    </table>

</body>
</html>`
)
