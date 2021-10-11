# October 11 2021

This is interesting:

> We're using HTTPS-based CDN audio file retrieval, similar to the official Web
> client or librespot-java, instead of the older, channel-based approach in
> librespot.

https://github.com/jpochyla/psst#thanks

Then look here:

https://github.com/librespot-org/librespot-java/blob/dev/lib/src/main/java/xyz/gianlu/librespot/dealer/ApiClient.java

First this:

~~~java
try (Response resp = send("GET", "/metadata/4/track/" + track.hexId(), null, null)) {
~~~

That is cool because it appears to be using HTTPS. `send` looks like this:

~~~java
Response resp = session.client().newCall(buildRequest(method, suffix, headers, body)).execute();
~~~

then `buildRequest` looks like this:

~~~java
request.addHeader("Authorization", "Bearer " + session.tokens().get("playlist-read"));
request.url(baseUrl + suffix);
~~~

then `baseUrl` looks like this:

~~~java
this.baseUrl = "https://" + ApResolver.getRandomSpclient();
~~~
