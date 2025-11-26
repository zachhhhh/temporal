# System Optimization & Performance

## Kernel Tuning (Linux)

Optimize the OS for high-throughput, low-latency workloads.

```bash
# /etc/sysctl.d/99-temporal.conf

# Network
net.core.somaxconn = 32768          # Increase backlog
net.ipv4.tcp_max_syn_backlog = 8192
net.ipv4.ip_local_port_range = 1024 65535
net.ipv4.tcp_tw_reuse = 1           # Reuse TIME_WAIT sockets

# TCP buffers (for high BDP)
net.core.rmem_max = 16777216
net.core.wmem_max = 16777216
net.ipv4.tcp_rmem = 4096 87380 16777216
net.ipv4.tcp_wmem = 4096 65536 16777216

# IO
fs.file-max = 2097152               # Open files limit
vm.swappiness = 10                  # Prefer RAM over swap
```

## Go Runtime Optimization

### Garbage Collection (GC)

Temporal is memory-intensive.

- `GOGC=100` (Default): Balanced.
- `GOGC=200`: Trade memory for CPU (less GC). Good for History service if RAM available.
- `GOMEMLIMIT`: Set to 90% of container limit to avoid OOM kills.

### Concurrency

- `GOMAXPROCS`: Auto-detect container limits (use `automaxprocs` library).

## Database Optimization (PostgreSQL)

### Connection Pooling

Use **PgBouncer** at transaction level.

- App -> PgBouncer (local sidecar) -> RDS
- Reduces connection overhead on Postgres.

### Query Tuning

- **Prepared Statements**: Always use.
- **JIT Compilation**: Enable for complex analytical queries.
- **Checkpointing**: Increase `max_wal_size` to reduce checkpoint frequency.

## Temporal-Specific Tuning

### History Shards

- Default: 512 shards.
- Cloud Scale: 4k, 8k, or 16k shards.
- **Rule**: ~1k workflows per shard active. Too many shards = CPU overhead. Too few = Lock contention.

### Cache Sizes

- `HistoryCacheMaxSize`: Increase to fit active working set in memory.
- Reduces DB IOPS significantly.

### Batching

- Enable `UseTransactionBatching` in history service.
- Batches DB writes for higher throughput.

## Load Balancer Tuning

### gRPC Tuning

- **HTTP/2 Keepalives**:
  - `PERMIT_WITHOUT_STREAM: true`
  - `MIN_TIME: 10s`
  - Prevents load balancers (ALB) from killing idle connections silently.

### Connection Balancing

- Use L7 load balancing (Envoy/Alb) to distribute gRPC _requests_, not just connections.
- Prevents "hot" backend pods.

## Benchmarking

### Standard Suite

Run `maru` (Temporal benchmark tool) weekly:

- **Throughput**: Max starts/sec.
- **Latency**: End-to-end workflow completion time.
- **Scalability**: Linear growth with added nodes?

### Profile Continuous (Pprof)

- Enable `net/http/pprof`.
- Continuous Profiling (e.g., Pyroscope) to find CPU hot paths in production.
