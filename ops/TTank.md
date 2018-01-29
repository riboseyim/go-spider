
```bash
initdb /Users/yanrui/Data/TTank
pg_ctl -D /Users/yanrui/Data/TTank -l logfile start
pg_ctl -D /Users/yanrui/Data/TTank -l logfile stop


ps -ef | grep postgre | grep -v 'grep' | awk '{print $2}'| xargs kill

```

```sql

DROP DATABASE ttank;

CREATE DATABASE ttank

GRANT CONNECT, TEMPORARY ON DATABASE ttank TO public;
GRANT ALL ON DATABASE ttank TO yanrui;
GRANT ALL ON DATABASE ttank TO ttank;



CREATE TABLE "public"."person" (
	"id" varchar(50) NOT NULL COLLATE "default",
	"name" varchar(50) COLLATE "default",
	"sex" varchar(10) COLLATE "default",
	"home" varchar(50) COLLATE "default",
	"position" varchar(500) COLLATE "default",
	"summary" text COLLATE "default",
	"resume" text COLLATE "default",
	"sourceurl" text COLLATE "default",
	"levelone" varchar(50) COLLATE "default",
	"leveltwo" varchar(50) COLLATE "default",
	"status" varchar(50) COLLATE "default",
	"ethnic" varchar(100) COLLATE "default",
	"birthday" varchar(100) COLLATE "default",
	"workday" varchar(100) COLLATE "default",
	"partyday" varchar(100) COLLATE "default",
	"education" varchar(200) COLLATE "default"
)
WITH (OIDS=FALSE);
GRANT ALL ON table person TO tianguan;

create table person_education (
Id       varchar(50),
record        varchar(10)
);

create table data_source_list (
Id         varchar(50) PRIMARY KEY NOT NULL,
itemcode   varchar(50),
title      varchar(100),
routepath  text,
status     varchar(50)
);

create table data_source_raw(
Id         varchar(50) PRIMARY KEY NOT NULL,
itemcode   varchar(50),
title      varchar(100),
rawvalue   varchar(100),
status     varchar(50)
);
GRANT ALL ON table data_source_raw TO tianguan;
```
