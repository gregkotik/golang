
Project consists of 2 services. One service parses the CSV file.
While reading the file, it calls a second service that either creates a new record in a database(Map), or updates the existing one.

The end result is a Map containing client records, representing the latest version found in the CSV. 
gRPC is being used for the communication between the 2 services
The phone numbers converted to the same format.
inclusing the country code (+44) to each one. 
