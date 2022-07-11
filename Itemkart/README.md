Rest Full Service List
----------------------------------
1. Get details of customer
    Request API URL  : http://localhost:50080/customer/1009
    Request API Type : GET

    Response : {
    "id": 1001,
    "name": "shri",
    "email": "shri@gmail.com",
    "address": "noida"
}
2. Create new customer
    Request API URL  : http://localhost:50080/customer
    Request API Type : POST
    Request API Body : {"id":1001,"name":"Shri","email":"shri@gmail.com","address":"Noida"}

    Response :{
    "message": "Customer successfully created",
    "status": "success"
}

3. Update customer record
    Request API URL  : http://localhost:50080/customer/1002
    Request API Type : POST
    Request API Body : {"name":"Boby","email":"boby@gmail.com","address":"Noida"}

    Response : {
    "message": "Customer record successfully updated",
    "status": "success"
}

4. Delete customer record
    Request API URL  : http://localhost:50080/customer/1009
    Request API Type : DELETE

    Response: {
    "message": "1 customer record deleted",
    "status": "success"
    }
----------------------------------
5. Get details of product
    Request API URL  : http://localhost:50080/product/101
    Request API Type : GET

    Response : {
    "id": 101,
    "name": "Mobile",
    "category": "redme",
    "description": "readme 9 pro"
}
6. Create new customer
    Request API URL  : http://localhost:50080/product
    Request API Type : POST
    Request API Body : {"id":101,"name":"Mobile","category":"redme","description":"redme 9 pro"}

    Response :{
    "message": "Product successfully created",
    "status": "success"
}

7. Update customer record
    Request API URL  : http://localhost:50080/Product/102
    Request API Type : POST
    Request API Body : {"name":"Mobile","category":"moto","description":"G3"}

    Response : {
    "message": "Product record successfully updated",
    "status": "success"
}

8. Delete customer record
    Request API URL  : http://localhost:50080/product/109
    Request API Type : DELETE

    Response: {
    "message": "1 product record deleted",
    "status": "success"
    }

