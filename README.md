# labclimate
##API Endpoints
| Method | Endpoint                           | Usage                                          |
|--------|------------------------------------|------------------------------------------------|
| GET    | /get_config                        | Gets the client configuration from the server. |
| POST   | /submit_clim                       | Submits climate data to the server.            |
| POST   | /nodes?id={id}&location={location} | Adds a new node to the server.                 |
| DELETE | /nodes?id={id}                     | Deletes a node from the server.                |
| PUT    | /nodes?id={id}&location={location} | Updates the location of a node on the server.  |
