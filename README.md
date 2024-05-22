# Rover example

Example scenario for rover dev using high cpu:

- Start the api
   - `cd api && go run main.go`
- Start rover dev
   - `rover dev --supergraph-config supergraph.yaml --router-config router.yaml`
- Open Activity Monitor, find `rover` process, note cpu usage
- Invoke the `echo` query for sanity (optional)
- Stop the api
   - Subgraph is removed
   - Note cpu usage rises to ~100%
- Start the api again
   - Subgraph is added back
   - Note cpu usage remains high