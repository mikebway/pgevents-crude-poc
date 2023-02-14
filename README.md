# A Crude Postgres Event Arc Audit Log Event Handler

As described in [Derivative Work](COPYRIGHT.md), the code this very simple Cloud Run container repository was
derived directly from the the Google Cloud Platform [golang-samples](https://github.com/GoogleCloudPlatform/golang-samples).

Thw original code has been modified slightly to log not Cloud Storage event information, but Postgres, `pgAudit`,
data change audit events.

The project does nothing terribly useful, just demonstrates that Postgres database record inserts and updates
can give rise to Pub/Sub events.

For more information on how to set up the larger context for which this code is just a final element, see  
[Enable Data Access audit logs](https://cloud.google.com/logging/docs/audit/configure-data-access),
[Audit for PostgreSQL using pgAudit](https://cloud.google.com/sql/docs/postgres/pg-audit), and
[Receive a Cloud Audit Logs event](https://cloud.google.com/eventarc/docs/run/cal).
