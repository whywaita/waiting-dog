# waiting-dog

:dog: waiting your request.

# URL

```
$ curl https://waiting-dog.herokuapp.com/  # health check
{"health": "ok"}

$ curl https://waiting-dog.herokuapp.com/wait/1 
# wait 1 second ...
{"message": "bow-wow!"}

$ curl https://waiting-dog.herokuapp.com/wait/100
# wait 100 second ...
{"message": "bow-wow!"}

$ curl https://waiting-dog.herokuapp.com/wait/random
# wait random time (1 - 10 second)
{"message": "bow-wow!"}
```
