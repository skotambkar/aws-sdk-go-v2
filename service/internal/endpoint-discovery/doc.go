/*
Endpoint discovery is a feature of AWS SDK for Go V2 that allows SDK
to fetch a valid endpoint to serve an API request. Discovered endpoints
are stored in an internal thread-safe cache to reduce the number of calls
made to fetch the endpoint.

Endpoint discovery uses two types of keys to store endpoint:
A. Customer specific key
B. Resource specific key

Each discovered endpoint has a TTL associated to it, and are evicted from
cache once expired. This feature can be turned on by setting
`AWS_ENABLE_ENDPOINT_DISCOVERY` env variable to true. By default, the feature
is set to AUTO - indicating operations that require endpoint discovery always
use it. To completely turn off the feature, one should set the value as false.
Similar configurations apply for shared config file where key is
`endpoint_discovery_enabled`.

*/
package endpointdiscovery
