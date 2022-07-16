package main

//定义规则
const rule1 = `
	rule "MedicalType-住院" "MedicalType为住院"  salience 2
	begin
		print(1)
		if mr.MedicalType != cond.MedicalType {
			return
		} 
		//if mr.ContainsAnyItem(cond.ItemsCode){
		//	rst.MatchSuccess = true
		//	rst.Rules = appendStr(rst.Rules, @name)
		//}
	end

	rule "MedicalRecord2"  "MedicalRecord ItemCode包含Condition Items的任何一个" salience 1
	begin
		print(2)
		if par1.MedicalType != par2.MedicalType {
			return
		} 
		//if mr.ContainsAnyItem(cond.ItemsCode){
		//	rst.MatchSuccess = true
		//	rst.Rules = appendStr(rst.Rules, @name)
		//} 
	end

	rule "MedicalRecordFor3"  "MedicalRecord ItemCode包含Condition Items的任何一个" salience 0
	begin
		print(3)
		if mr.MedicalType != cond.MedicalType {
			return
		} 
		mrItems := mr.Items
		condItems := cond.ItemsCode
		forRange idx := mrItems {
			forRange i := cond.ItemsCode {
				elem := mrItems[idx]
				if elem.Code == condItems[i] {
					rst.MatchSuccess = true
					rst.Rules = appendStr(rst.Rules, @name)
				}
			}
		}
	end

	rule "4"  "" salience 1
	begin
		print(4)
	end

	rule "5"  "" salience 1
	begin
		print(5)
	end

	rule "6"  "" salience 1
	begin
		print(6)
	end

	rule "7"  "" salience 1
	begin
		print(7)
	end

	rule "8"  "" salience 1
	begin
		print(8)
	end

	rule "9"  "" salience 1
	begin
		print(9)
	end

	rule "10"  "" salience 1
	begin
		print(10)
	end
`
