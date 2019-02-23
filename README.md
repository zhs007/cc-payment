# cc-payment

[![Build Status](https://travis-ci.org/zhs007/cc-payment.svg?branch=master)](https://travis-ci.org/zhs007/cc-payment)

This is a coding challenge for payment.  
You can start the unit test by executing [``startunittest.sh``](https://github.com/zhs007/cc-payment/startunittest.sh).  
The unit test environment is in [``Dockerfile.unittest``](https://github.com/zhs007/cc-payment/Dockerfile.unittest).  

This project uses [Gin](https://github.com/gin-gonic/gin), [Protocol Buffers](https://developers.google.com/protocol-buffers/), and [zap](https://github.com/uber-go/zap).  
The database is ``mysql 5`` .  

The public structure definition uses [Protocol Buffers](https://developers.google.com/protocol-buffers/), see [proto/payment.proto](https://github.com/zhs007/cc-payment/proto/payment.proto).  
The front end can use it.  

The database script is in [``sql/ccpayment.sql``](https://github.com/zhs007/cc-payment/sql/ccpayment.sql).  

You can build and deploy with [``Dockerfile``](https://github.com/zhs007/cc-payment/Dockerfile).  
You can build and deploy with [``builddocker.sh``](https://github.com/zhs007/cc-payment/builddocker.sh) and [``startdocker.sh``](https://github.com/zhs007/cc-payment/startdocker.sh).    
You need to configure ``cfg/config.yaml`` first.  
We did not support https in the project, you can configure it in nginx.  

About the numeric type of currency, we use int64 instead of floating point numbers.  
For example, USD, 100 units means 1 US dollar.  

There is no calculation fee for this version.  

You can see the payment logic in [``model/paymentdb.go``](https://github.com/zhs007/cc-payment/model/paymentdb.go).    

[Payment API Document](https://app.swaggerhub.com/apis-docs/zhs007/cc-payment/1.0.0)