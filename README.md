labclimate
==========

API Endpoints
---------------

| Method | Endpoint         | Usage                                          |
|--------|------------------|------------------------------------------------|
| GET    | /config          | Gets the client configuration from the server. |
| POST   | /data            | Submits climate data to the server.            |
| POST   | /nodes           | Adds a new node to the server.                 |
| DELETE | /nodes/{id}      | Deletes a node from the server.                |
| PUT    | /nodes/{id}      | Updates the location of a node on the server.  |

### submit_clim POST body:

node_id: {node id},
temp: {temperature},
humid: {humidity},
air_qual: {air quality value},
time: {timestamp}

### nodes POST body:

node_id: {node id},
node_loc: {location}

### nodes/{id} PUT body:

node_loc: {location}

