# October 12 2021

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

So first, we need to get random client:

~~~
GET /?type=accesspoint HTTP/1.1
Host: apresolve.spotify.com
~~~

Result:

~~~json
{
  "accesspoint": [
    "guc3-accesspoint-a-351z.ap.spotify.com:4070",
    "guc3-accesspoint-a-zrnq.ap.spotify.com:443",
    "guc3-accesspoint-a-g2wl.ap.spotify.com:80",
    "guc3-accesspoint-a-j2k6.ap.spotify.com:4070",
    "guc3-accesspoint-a-bg8f.ap.spotify.com:443",
    "guc3-accesspoint-a-h4q0.ap.spotify.com:80",
    "gew1-accesspoint-a-grhm.ap.spotify.com:4070",
    "gew1-accesspoint-a-2gzw.ap.spotify.com:443",
    "gae2-accesspoint-a-279c.ap.spotify.com:80"
  ]
}
~~~

Then take radix 62 track ID:

~~~
7gTsnFUJMoBIMIzFRUi8to
~~~

convert to radix 16:

~~~
eef38251727f46c28eed9284b288024e
~~~

For the connection to work, we need ClientHello:

https://github.com/librespot-org/librespot-java/blob/353c9db0/lib/src/main/java/xyz/gianlu/librespot/core/Session.java#L212

This site says `apresolve.spotify.com` isnt used any more:

https://github.com/ahixon/spotify-mitm-proxy#ap-resolve-server

and that `ap.spotify.com` is used.
