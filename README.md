## concurrency

# fundamentals 
- every program you execute is a Linux process that can run on your machine's CPU (core)
- not all tasks in your program will require your machines core
    - database fetches is the easiest example. it requires a network call (through your computer's interface) and the core your DB is running on 
- during these tasks, your core is free to perform other tasks 
- concurrency is the term for tasks that can run without your CPU core and can therefore be fired "in parallel" to maximize the usage of your core (speed up your program)
- parallelism would require another core to run tasks on. it cannot be achieved with only 1 core 
