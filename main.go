package main
import "fmt"
import "github.com/olekukonko/tablewriter"
import "os"
import "strconv"

func main() {


	var a1 float64
	var a2 float64

	fmt.Println("请输入正股价: ")
	fmt.Scanln(&a1)
	
	fmt.Println("请输入转股价: ")
	fmt.Scanln(&a2)


	probability := []map[string]interface{}{
		map[string]interface{}{"v0":"极度乐观","v1":0.2,"v2":0.05},
		map[string]interface{}{"v0":"一般乐观","v1":0.1,"v2":0.1},
		map[string]interface{}{"v0":"较为乐观","v1":0.05,"v2":0.2},
		map[string]interface{}{"v0":"正常","v1":0.00,"v2":0.3},
		map[string]interface{}{"v0":"较为悲观","v1":-0.05,"v2":0.2},
		map[string]interface{}{"v0":"一般悲观","v1":-0.1,"v2":0.1},
		map[string]interface{}{"v0":"极度悲观","v1":-0.2,"v2":0.05},
	}

	data := make([][]string,7)

	expect := 0.00
	gain := 0.00 


	for _, v := range probability {
		v["v3"] = (1 + v["v1"].(float64)) * a1
		v["v3"],_ = strconv.ParseFloat(fmt.Sprintf("%.2f", v["v3"]), 64)
		v["v4"] = a2
		v["v5"] = (100.00 / v["v4"].(float64)) * v["v3"].(float64);
		v["v6"] = v["v5"].(float64) * 1.1
		v["v7"] = v["v6"].(float64) - 100
		data = append(data,[]string{ 
			v["v0"].(string) ,
			strconv.FormatFloat(v["v1"].(float64) *100, 'f', 0, 32)+"%" ,
			strconv.FormatFloat(v["v2"].(float64) *100, 'f', 0, 32)+"%",
			strconv.FormatFloat(v["v3"].(float64) , 'f', 2, 32),
			strconv.FormatFloat(v["v4"].(float64) , 'f', 2, 32),
			strconv.FormatFloat(v["v5"].(float64) , 'f', 2, 32),
			strconv.FormatFloat(v["v6"].(float64) , 'f', 2, 32),
			strconv.FormatFloat(v["v7"].(float64) , 'f', 2, 32),
		})

		expect += v["v2"].(float64) * v["v7"].(float64)

		if(v["v7"].(float64) > 0){
			gain += v["v2"].(float64)
		}
	}
	
	expect_str := "预期收益："+fmt.Sprintf("%.2f", expect) +"/张"
	gain_str := ""
	gain_str +=  fmt.Sprintf("%.0f", gain*100) + "%|" 
	resultColor := tablewriter.Colors{tablewriter.FgHiRedColor}
	if(gain >= 0.65 ){
		gain_str += "建议买入"
		resultColor = tablewriter.Colors{tablewriter.FgHiGreenColor}
	}else{
		gain_str += "建议放弃" 
	}


	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"模拟行情", "涨跌幅度", "出现概率","正股价","转股价","转股价值","转债价格","预期收益"})
	table.SetFooter([]string{"", "","","","","", expect_str, gain_str})

	table.AppendBulk(data)
	table.SetFooterColor(tablewriter.Colors{}, tablewriter.Colors{}, tablewriter.Colors{}, tablewriter.Colors{}, tablewriter.Colors{}, tablewriter.Colors{},
		tablewriter.Colors{tablewriter.Bold},
		resultColor)
	table.Render() // Send output


	

}
