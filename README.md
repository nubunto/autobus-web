# Autobus Web

This is a Goa project. That means that the API documentation is always 100% up-to-date.
That also means that parts of this project may present a cognitive overhead.

# Architecture

- For web stuff, we have [Goa](https://goa.design). We let Goa take care of things. We don't step at Goa's toe, and it doesn't step in ours.
- The Autobus service that actually listens to the thoughts and feelings of the GPSs is in [this repository](https://github.com/nubunto/autobus).

We let Goa handle HTTP and JSON. The responses are neatly described in the generated Swagger docs.

Our handlers fetch data from a PostgreSQL database, in order to tell the user what the fuck is going on.

If you wanna take a peak at the service that parses and persists the data in a PostgreSQL database, take a look at the [Autobus Platform](https://github.com/nubunto/autobus-platform).
