# hcaptcha-solver-api

An extension for the [hcaptcha-solver-go](https://github.com/JustTalDevelops/hcaptcha-solver-go) package I made,
allowing services that use other languages to use the solver.
It uses Fiber to run an HTTP server on port 80, 
and then handles captcha solving requests.

# Usage

If you would like to use Google's Vision API to determine answers,
set the `GOOGLE_APPLICATION_CREDENTIALS` environment variable on your PC.
If not set, it will default to randomly guessing.

After setting the environment variable, do the following:

```
go install github.com/justtaldevelops/hcaptcha-solver-api
hcaptcha-solver-api
```

If all goes well, you should see the following:
```
hCaptcha solver API is now running on port 80!
```

The console will now display information about ongoing requests.

# Configuration
You can configure certain options of the API to your liking.
Currently, you can change the `Authorization` header, the port,
and the solution deadline.

# API Usage

Solving tasks are handled with a `POST` request to `example.org/solve`.
Make sure your `Authorization` header matches the one in your config.
Your request body should be the following:
```
{"site_url": "example.com"}
```

For additional options, refer to [this portion](https://github.com/JustTalDevelops/hcaptcha-solver-go/blob/5ec4aaabf52cc71e3e906fe7fba2a3555ef9f8f6/solver.go#L42) of the original package.
An example of additional options in JSON are as follows:
```
{"site_url": "example.com", "options": {"site_key": "8fc2ffd0-a7a5-4345-a357-97d389a9a635"}}
```

After sending your request, a response will be returned in JSON which will be something like:
```
{"captcha_code": "P0_eyJ0eXAiOiJKV1Q..."}
```

You can then use this captcha code in your application.