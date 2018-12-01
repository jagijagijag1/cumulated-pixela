# cumulated-pixela
On your [Pixela](https://pixe.la/) graph, the cumulated-pixela just read yesterday's pixel (quantity) and record it to today's pixel.

## Project setup
### Requirements
- Go environment 1.11.2
- serverless framework 1.34.1

### compile & deploy
```bash
git clone https://github.com/jagijagijag1/cumulated-pixela
cd cumulated-pixela
```

Describe your pixela info to `environment` clause on `serverless.yml`.
If you want, change the time for periodic invoking of Lambda function (default: 16:00 UST = 01:00 JST) in `schedule`.

```yaml:serverless.yml
...
functions:
  cumulated-pixela:
    handler: bin/cumulated-pixela
    events:
      - schedule: cron(0 16 * * ? *)
    # you need to fill the followings with your own
    environment:
      TZ: Asia/Tokyo
      PIXELA_USER: <user-id>
      PIXELA_TOKEN: <your-token>
      PIXELA_GRAPH: <your-graph-id-1>
    timeout: 10
```

Then, run the following.

```bash
make
sls deploy
```
