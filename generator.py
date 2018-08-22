from peewee import *;

db = SqliteDatabase('the.db');

class Block(Model):
    height = IntegerField();
    time = CharField();
    txs_n = IntegerField();
    inner_txs_n = IntegerField();
    txs = CharField();
     
    class Meta:
        database = db;

db.connect();
db.create_tables([Block]);

def insert():
    Block.create(height = 0, time = 'now', txs_n = 0, inner_txs_n = 1, txs = 'david');
    print("Insert Succeed!")

def read():
    print("total query: [", Block.select().count(), "]")
    for b in Block.select():
        print(b.id, "time:", b.time)

arg = input("choose function: (1) insert data  (2) read data :  ");
if arg == str(1):
    insert();
else:
    read();
