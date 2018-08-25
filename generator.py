from peewee import *;

db = SqliteDatabase('the.fox');

class Block(Model):
    difficulty = IntegerField();
    gas_limit = IntegerField();
    gas_used = IntegerField();
    hash = CharField();
    number = IntegerField();
    size = IntegerField();
    timestamp = IntegerField();
    total_difficulty = IntegerField();
    txs_n = IntegerField();
    finished = IntegerField();
     
    class Meta:
        database = db;

db.connect();
db.create_tables([Block]);

def insert():
    Block.create(
        difficulty = 1,
        gas_limit = 1,
        gas_used = 1,
        hash = 'h',
        number = 1,
        size = 1,
        timestamp = 1,
        total_difficulty = 1,
        txs_n = 1,
        finished = 1,
    );
    print("Insert Succeed!")

def read():
    print("total query: [", Block.select().count(), "]")
    for b in Block.select():
        print(b.id, "time:", b.timestamp)

arg = input("choose function: (1) insert data  (2) read data :  ");
if arg == str(1):
    insert();
else:
    read();
