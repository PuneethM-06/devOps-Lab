# Day 10 - Interview Questions

## Ports

1. What is the difference between an IP address and a port?

2. What does it mean when a process is "listening" on a port?

3. Can two applications listen on the same port simultaneously? Why or why not?

---

## TCP

4. What is TCP and what problems does it solve?

5. How does TCP ensure reliable delivery of data?

6. What is the difference between ordered delivery and reliable delivery in TCP?

---

## TCP Three-Way Handshake

7. Explain the TCP three-way handshake step by step.

8. Why is a three-way handshake required before data transmission?

9. A client sends a SYN packet but never receives a SYN-ACK. What are some possible causes?

---

## UDP

10. What is UDP and how is it different from TCP?

11. Why does DNS commonly use UDP instead of TCP?

12. Give two real-world use cases where UDP is preferred over TCP.

---

## Ephemeral Ports

13. What is an ephemeral port and why is it required?

14. In the connection below, identify the client port and server port:

192.168.1.10:55000 → 10.0.0.5:443

---

## Troubleshooting & Connection States

15. Explain the meaning of the following TCP states:

* LISTEN
* ESTABLISHED
* TIME_WAIT
* CLOSE_WAIT

16. Which command would you use on Linux to:

* View listening ports?
* View active TCP connections?

---

## End-to-End Networking Flow

17. Explain what happens when a user enters https://google.com in a browser until the webpage is displayed.
