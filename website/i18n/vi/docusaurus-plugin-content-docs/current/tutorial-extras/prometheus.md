---
sidebar_position: 1
---

# Using Prometheus with Veles

Using Prometheus with Veles is really easy as Veles exposes a compatible metrics endpoint by default.

To use the metrics simply add a scrape-job to Prometheus to pull the metrics from `http://[your-server-here]:8123/metrics`.

And that's it, you now have a few relevant metrics to monitor the health of your Veles! 🎉