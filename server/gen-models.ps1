gen --sqltype=mysql `
	--connstr="daigou:daigou123!!@tcp(127.0.0.1:3306)/daigou_test?parseTime=True&loc=Local&charset=utf8mb4&collation=utf8mb4_bin" `
	--database=daigou  `
	--json `
	--json-fmt=snake `
	--gorm `
	--model=dg_model `
	--out=./model `
	--overwrite `
	--guregu `
	--verbose `
	--templateDir=./template/db `
	--mapping=./template/db/mapping.json `
	--model_naming="{{FmtFieldName .}}" `
	--field_naming="{{FmtFieldName (stringifyFirstChar .) }}" `
	--file_naming="{{.}}" `
	--windows `
	--module=daigou-admin/model `
	# --generate-dao `
	# --dao=dao `
	# --save=./model/daigoudb/template `
# --name_test=orders_goods `