# Interview.md — Day 84

## 1. What is Alertmanager?

Alertmanager is a component in the Prometheus ecosystem that receives firing alerts from Prometheus and manages how those alerts are handled.

It is responsible for:

- Grouping
- Routing
- Deduplication
- Silencing
- Inhibition
- Sending notifications to receivers such as Slack or email

---

## 2. What is the difference between Prometheus and Alertmanager?

Prometheus collects metrics and evaluates alert rules.

When an alert condition becomes true and remains true for the configured `for` duration, Prometheus changes the alert state to **Firing** and sends it to Alertmanager.

Alertmanager then manages the alert by grouping, deduplicating, silencing, inhibiting, routing, and sending notifications.

> Prometheus detects the problem.

> Alertmanager manages and routes the alert.

---

## 3. Explain the alert lifecycle.

The alert lifecycle is:

    Inactive
        ↓
    Condition becomes True
        ↓
    Pending
        ↓
    Condition remains true for `for` duration
        ↓
    Firing
        ↓
    Sent to Alertmanager

When the condition becomes false, the alert is resolved.

---

## 4. What is the difference between `rate(metric[5m])` and `for: 5m`?

`rate(metric[5m])` uses the previous 5 minutes of metric data to calculate the per-second average rate.

`for: 5m` means the alert condition must remain true continuously for 5 minutes before becoming **Firing**.

---

## 5. What is grouping in Alertmanager?

Grouping combines related alerts into fewer notifications.

For example, if multiple pods of the same service trigger similar alerts, Alertmanager can group them and send one meaningful notification instead of sending a separate notification for each pod.

---

## 6. What is routing in Alertmanager?

Routing determines where an alert should be sent.

Alertmanager can use labels such as `team="payments"`, `severity="critical"`, or `service="database"` to route alerts to the correct receiver, such as a specific Slack channel, email address, or on-call team.

---

## 7. What is deduplication?

Deduplication prevents unnecessary repeated notifications for the same alert.

If an alert is already firing, Alertmanager recognizes repeated updates for that alert and avoids flooding the notification channel with duplicate messages.

---

## 8. What is the difference between silencing and inhibition?

**Silencing** is manually configured to temporarily suppress notifications for alerts matching specific conditions. It is commonly used during planned maintenance.

**Inhibition** automatically suppresses related or lower-priority alerts when another more important alert is already firing.

Example:

    DatabaseDown
        ↓
    Automatically suppress:
    - DatabaseConnectionError
    - ApplicationError

---

## 9. Explain the flow from an error spike to a Slack notification.

The application exposes metrics, and Prometheus scrapes those metrics.

Prometheus evaluates an alert rule. When the error condition becomes true, the alert enters the **Pending** state.

If the condition remains true for the configured `for` duration, the alert becomes **Firing**.

Prometheus sends the firing alert to Alertmanager.

Alertmanager then applies grouping, deduplication, silencing, inhibition, and routing.

Finally, the alert is sent to the configured receiver, such as Slack.

The complete flow is:

    Application
        ↓
    Metrics
        ↓
    Prometheus
        ↓
    Alert Rule
        ↓
    Pending
        ↓
    Firing
        ↓
    Alertmanager
        ↓
    Group / Deduplicate / Silence / Inhibit
        ↓
    Route
        ↓
    Receiver
        ↓
    Slack Notification

---

## 10. What is the role of a Slack webhook in Alertmanager?

A Slack webhook provides the endpoint that Alertmanager uses to send notifications to Slack.

The Slack receiver configuration contains the webhook URL and notification settings, such as the destination channel and whether resolved notifications should also be sent.