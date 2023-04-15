# scaffold

## What is scaffold?

Scaffold is the meta-framework that is used to build and manage, and facilitate interactions with robots 
developed around the [Loppu framework](https://github.com/jt05610/loppu). 

## Why did I make it?
The conceptual model of a loppu robot is a network of components that individually do one thing well, and together can accomplish complex tasks.
To achieve this, loppu is built around standardized communications protocols that are optimized for interaction with scalable networks of both
hardware and software nodes.

Given the standardized protocols, most of the device firmware can be automatically generated, and only the device-specific functionality needs to be worried about. 
So, I made this meta-framework to maintain standardization and save myself countless hours of time writing boilerplate code, while also making it easier for others 
to use the things I come up with.

## Inspiration
My main intention with Loppu is controlling scalable networks of as yet unknown devices that interact within a laboratory environment. I therefore approached this
with heavy inspiration from modular web framework such as Django and Ruby on Rails, and less inspiration from robotics frameworks such as ROS. 

## How does I use it to make robots?

### Initializing a loppu robot

From directory that you plan to keep your robots in, run `scaffold init -n <device_name>`. This will create a directory with the standardized directory structure.

### Creating nodes for your robot

The robot's functionality is brought about from composition of nodes that each do one thing well. Nodes can be software or hardware, and loppu exists to ensure 
communication between both kinds of nodes is seamless and optimized for laboratory-based data needs. We can both create new nodes and reuse one that have already made.

#### Creating a hardware node

##### 1. Generating the structure
To create a new hardware node, run `scaffold node new hardware <your_node_name>`. By default, your node is generated in `~/scaffold/nodes/<your_node_name>`

##### 2. Describing the functionality
To describe the functionality of your node, open `~/scaffold/nodes/<your_node_name>/node.yaml` with your favorite text editor. It should look like this:

```
meta:
    node: <your_node_name>
    author: <your_date>
    address: <modbus address>
    date: <date>
tables:
    coils:
        - name: coil_1
          desc: Coils are binary read/write registers. They are used to execute on/off functions on the target device
          params:
            - value:
                type: int
                desc: Writing 0 will write 0 to the device. Writing anything else will write 1.
        - name: coil_2
          desc: They can also be parameter-less if you want
    discrete_inputs:
        - name: discrete_input_1
          desc: Discrete inputs are binary read only registers
        - name: discrete_input_2
          desc: Add a new one like this
    holding_registers:
        - name: holding_register_1
          desc: Holding registers are 16-bit read/write registers. They are used to set variables on the device, and can execute a if desired
          params:
            - value:
                type: int
                desc: Whatever you write will be converted to a uint16
        - name: holding_register_2
          desc: Add a new one like this
          params:
            - value:
                type: int
                desc: You technically don't need to include a parameter but you probably should
    input_registers:
        - name: input_register_1
          desc: Input registers are 16-bit read only registers
        - name: input_register_2
          desc: Add a new one like this

```



