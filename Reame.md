# Introduction to Protocol Buffer..
# Introduction to Protocol Buffers

Protocol Buffers, also known as Protobuf, is a language-agnostic binary serialization format developed by Google. It is designed to efficiently serialize structured data for communication between different systems.

## Why use Protocol Buffers?

- **Efficiency**: Protocol Buffers use a binary format, which makes them more compact and faster to serialize and deserialize compared to other data interchange formats like JSON or XML.
- **Language-agnostic**: Protocol Buffers support multiple programming languages, including C++, Java, Python, and more. This allows you to define your data structures once and generate code for different languages.
- **Forward and backward compatibility**: Protocol Buffers provide built-in support for versioning and evolving your data schema without breaking existing clients or servers.
- **Schema definition**: Protocol Buffers use a schema definition language (`.proto` files) to define the structure of your data. This schema acts as a contract between the sender and receiver, ensuring data consistency.

## How Protocol Buffers work

1. Define your data structure using the Protocol Buffers schema definition language (`.proto` file).
2. Compile the schema using the Protocol Buffers compiler (`protoc`) to generate language-specific code.
3. Use the generated code in your application to serialize and deserialize data using Protocol Buffers.

## Example

Here's a simple example of a `.proto` file defining a `Person` message:
syntax = "proto3";

```
package example;

// Example of a person with Go proto including both service and message format

// The package name for the generated Go code
package example

// Person message represents a person with name and age
message Person {
    string name = 1; // Name of the person
    int32 age = 2; // Age of the person
}

// PersonService is a service that provides operations related to persons
service PersonService {
    // GetPerson is an RPC method that takes a PersonRequest and returns a PersonResponse
    rpc GetPerson(PersonRequest) returns (PersonResponse);
}

// PersonRequest message represents a request to get a person by their ID
message PersonRequest {
    string person_id = 1; // ID of the person
}

// PersonResponse message represents a response containing a person and a message
message PersonResponse {
    Person person = 1; // The person
    string message = 2; // A message related to the response
}
```
## You must give each field in your message definition a number between 1 and 536,870,911 with the following restrictions:

-> The given number must be unique among all fields for that message.
-> Field numbers 19,000 to 19,999 are reserved for the Protocol Buffers implementation. The protocol buffer compiler will complain if you use one of these reserved field numbers in your message.
--> You cannot use any previously reserved field numbers or any field numbers that have been allocated to extensions.
--> This number cannot be changed once your message type is in use because it identifies the field in the message wire format. “Changing” a field number is equivalent to deleting that field and creating a new field with the same type but a new number. See Deleting Fields for how to do this properly.

--> Field numbers should never be reused. Never take a field number out of the reserved list for reuse with a new field definition. See Consequences of Reusing Field Numbers.

--> You should use the field numbers 1 through 15 for the most-frequently-set fields. Lower field number values take less space in the wire format. For example, field numbers in the range 1 through 15 take one byte to encode. Field numbers in the range 16 through 2047 take two bytes. You can find out more about this in Protocol Buffer Encoding.

# Example 2. gRPC service definition of ProductInfo service using protocol buffers
```
syntax = "proto3"; 
package ecommerce; 

service ProductInfo { 
    rpc addProduct(Product) returns (ProductID); 
    rpc getProduct(ProductID) returns (Product); 
}

message Product { 
    string id = 1; 
    string name = 2;
    string description = 3;
}

message ProductID { 
    string value = 1;
}
```

## 1 The service definition begins with specifying the protocol buffer version (proto3) that we use.

## 2 Package names are used to prevent name clashes between protocol message types and also will be used to generate code.

## 3 Definition of the service interface of the service.

## 4 Remote method to add a product that returns the product ID as the response.

## 5 Remote method to get a product based on the product ID.

## 6 Definition of the message format/type of Product.

## 7 Field (name-value pair) that holds the product ID with unique field numbers that are used to identify your fields in the message binary format.

## 8 Definition of the message format/type of ProductID.
