# Doohan

This is gonna be some kind of timetracking software, written in Go, Revel and using Postgres.

It’s my first Go project, I have no idea what I’m doing, and I guess it shows. But it’s working.

Don’t run this code, don’t fork it, don’t use it in production. It’s less than Alpha.

## How to install

Didn’t you listen? Don’t run this!

Okay, whatever. You need Postgres running with a `doohan` user (password `doohan`) who is the owner of a `doohan` database. Then:

    go get github.com/revel/cmd/revel bitbucket.org/liamstask/goose/cmd/goose
    goose up
    revel run github.com/scy/doohan

Call http://localhost:9000/ to see a list of timetracking entries. To insert them, POST to http://localhost:9000/add and supply a `description` parameter, like:

    curl -v -d description='trying to get Doohan to work' http://localhost:9000/add

## Database stability

Extra special warning: Yes, Doohan has Goose to be able to upgrade database layouts. But: **During development, I won’t create migration files for every single poop that I’m doing. Instead, I’ll be updating the “Initial” file until I think it’s stable enough. Don’t insert production data. Really.**
