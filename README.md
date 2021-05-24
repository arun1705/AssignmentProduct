# AssignmentProduct

#Steps to run

Step 1: Start the Fabric test-network using below command.

cd fabric-samples
cd test-network
sudo ./network.sh up -ca -s couchdb

Step 2: Once network is up and running then create new channel using below command.
sudo ./network.sh createChannel -c testchannel


Step 3: Execute chaincode lifecycle shell file that performs all required lifecycle steps.
./chaincodeLifecycle.sh

Step 4: Download the modules using below command.
cd clientApp
npm install

Step 5: Run the below command to enroll the admin user.
cd clientApp
node enrolladmin.js

Step 6: Run the below command to enroll the admin user.
cd clientApp
node registerEnrollClientUser.js

Step 7: Run the below command to start the server
cd clientApp
node server.js

Step 8: Test the api's on postman
a)/api/addproduct/
Input for this API call.
{
"id": "1",
"productName": "product1",
"description": "something...",
"prize": "1000",
"colour": "blue"
}

b)/api/queryproductbyid/:id

c)/api/queryallproducts

d)/api/deleteproductbyid/
Input for this API call.
{
"id": "1"
}

