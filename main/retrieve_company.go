package main

import (
	
//	"fmt"
//	_ "github.com/lib/pq"
	_ "github.com/jinzhu/gorm/dialects/postgres"		
//	"github.com/jinzhu/gorm"
)
	 
func retrieve_company() {
	
	
	rows, err := db.Query("SELECT * FROM companydata;")

	checkErr(err, "Error on query DB")	

    for rows.Next() {
	 
            var c Company
            
            err = rows.Scan(&c.name, &c.number, &c.careOf, &c.poBox, &c.addressLine1, 
            	&c.addressLine2, &c.postTown, &c.county, &c.country, &c.postcode, 
	            &c.category, &c.status,  &c.countryOfOrigin,  &c.dissolution_date,  &c.incorporate_date,
	            &c.accounting_refDay,  &c.accounting_refMonth,  &c.account_nextDueDate,  &c.account_lastMadeUpdate,  &c.account_category,
	            &c.return_nextDueDate,  &c.return_lastMadeUpdate,  &c.num_MortChanges,  &c.num_MortOutstanding,  &c.num_MortPartSatisfied,
	            &c.num_MortSatisfied,  &c.siccode1,  &c.siccode2,  &c.siccode3,  &c.siccode4,
	            &c.num_genPartner, &c.num_limPartner, &c.uri, &c.pn1_condate, &c.pn1_companydate, 
	            &c.pn2_condate, &c.pn2_companydate, &c.pn3_condate, &c.pn3_companydate, &c.pn4_condate,
	            &c.pn4_companydate,&c.pn5_condate, &c.pn5_companydate, &c.pn6_condate, &c.pn6_companydate, 
	            &c.pn7_condate, &c.pn7_companydate, &c.pn8_condate, &c.pn8_companydate, &c.pn9_condate, 
	            &c.pn9_companydate, &c.pn10_condate, &c.pn10_companydate, &c.conf_stmtNextDueDate, &c.conf_stmtLastMadeUpdate)
			checkErr(err, "Read company data rows,")
			
//			fmt.Printf("%+v\n", c)
    }

}

