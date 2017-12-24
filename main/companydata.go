package main

import (
	"database/sql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"

)

type Company struct {

	name 					sql.NullString
	number 					string 
	careOf 					sql.NullString 
	poBox 					sql.NullString  
	addressLine1 			sql.NullString // 5
	
	addressLine2 			sql.NullString  
	postTown 				sql.NullString  
	county 					sql.NullString  
	country 				sql.NullString  
	postcode 				sql.NullString // 10
	
	category 				string 
	status 					string 
	countryOfOrigin 		string 
	dissolution_date 		sql.NullString
	incorporate_date 		sql.NullString // 15
	 
	accounting_refDay 		sql.NullInt64
	accounting_refMonth 	sql.NullInt64
	account_nextDueDate 	sql.NullString
	account_lastMadeUpdate 	sql.NullString 
	account_category 		sql.NullString // 20
	
	return_nextDueDate 		sql.NullString
	return_lastMadeUpdate 	sql.NullString 
	num_MortChanges 		int64 
	num_MortOutstanding 	int64
	num_MortPartSatisfied 	int64 // 25
	
	num_MortSatisfied 		int64
	siccode1 				sql.NullString 
	siccode2 				sql.NullString
	siccode3 				sql.NullString
	siccode4 				sql.NullString // 30
	
	num_genPartner 			int
	num_limPartner 			int
	uri 					string 
	pn1_condate 			sql.NullString 
	pn1_companydate 		sql.NullString // 35
	
	pn2_condate 			sql.NullString 
	pn2_companydate 		sql.NullString
	pn3_condate 			sql.NullString 
	pn3_companydate 		sql.NullString
	pn4_condate 			sql.NullString // 40
	
	pn4_companydate 		sql.NullString
	pn5_condate 			sql.NullString 
	pn5_companydate 		sql.NullString
	pn6_condate 			sql.NullString 
	pn6_companydate 		sql.NullString // 45
	
	pn7_condate 			sql.NullString 
	pn7_companydate 		sql.NullString
	pn8_condate 			sql.NullString 
	pn8_companydate 		sql.NullString
	pn9_condate 			sql.NullString // 50
	
	pn9_companydate 		sql.NullString
	pn10_condate 			sql.NullString 
	pn10_companydate 		sql.NullString
	conf_stmtNextDueDate 	sql.NullString 
	conf_stmtLastMadeUpdate sql.NullString // 55
}
