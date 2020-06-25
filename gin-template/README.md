# Gin-API

![Build Status](https://github.com/airdb/passport/workflows/Go/badge.svg)


# Gin Web API

Workflow as follow
```
                    +----------+
                    |          |
                    | Webrouter|
                    |          |
                    +----------+
                         ^
                         |
                         v
                    +----------+
                    |          |
                    | WebHander|
                    |          |
                    +----------+
                         ^
                         |
                         v
                    +----------+
                    |          |
                    |   VO     |
                    |          |
                    +----------+
                         ^
                         |
                         v
+----------+                          +----------+
|          |                          |          |
|    PO    |       <---(or)--- >      |    BO    |
|          |                          |          |
+----------+                          +----------+
     ^
     |
     v
+---------+
|         |
|   DB    |
|         |
+---------+
```

## 

WebRouter: Match individual URL components to a Controller and a method defined

WebHandler: Handle gin context, query args and post body.

VO: Value Object

BO: Business Object

PO: Persistant Object

