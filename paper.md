---
title: 'Loppu: a laboratory robotics framework for getting research done'
tags:

- robotics
- C
- Go
- Python
- concurrency
- networking
- embedded
- biotechnology

  authors:
- name: Jonathan Taylor
  orcid: 0000-0001-5281-6990
  equal-contrib: true
  affiliation: "1"  

  affiliations:
- name: Skaggs School of Pharmacy and Pharmaceutical Sciences, University of Colorado Anschutz Medical Campus, Aurora, Colorado, USA.
  index: 1
  date: 14 May 2023
  bibliography: paper.bib

# Summary

Writing software for robots is difficult `quigley2009ros`. To address this,
excellent frameworks such as Robot Operating System (ROS) `quigley2009ros` have
been developed to help robotics developers build better robots and write less
code. While these tools are indispensable to robotics researchers, they
remain difficult use for researchers who are not robotics researchers but
would benefit from incorporating robotics into their workflows. While there
are many scientific robots and instruments available to purchase for research
tasks such as those in analytical chemistry and pharmaceutical screening, there
are simply too many esoteric and one-time experiments to have a robot available
for every possible scientific need. `Loppu` is a robotics framework intended 
to help researchers build their own tools that they would otherwise not be 
able to purchase. It is not intended to take the place of any existing 
robotics frameworks or scientific robots, it is intended to make developing 
and controlling autonomous systems around existing and novel robots easier.

# Statement of need

`Loppu` is a Go and C framework intended to build robots to automate and
control laboratory processes. The philosophy of `Loppu` is heavily inspired
by the Unix philosophy in that robots are intended to be built of multiple
small nodes that do one thing well, and complexity arises through composition.
`Loppu` provides a command line interface for building, maintaining, and
reusing modular robots that are made of both software and hardware nodes.
Software nodes are intended to interact with data -- databases, AI model
training, automated analysis, etc. Hardware nodes are intended to
do the physical work -- control stepper motors, read sensors, etc. Hardware
nodes support Modbus over Serial and GCode protocols. Modbus was decided as
the most efficient and reliable protocol for small hardware nodes, and GCode
is supported so that existing devices such as 3D printers and autosamplers can 
easily be hacked into XYZ positioning systems that interact within a larger 
system. Additionally, a meta-framework is provided to automatically generate 
as much code as possible so that researchers can focus on what their robot 
needs to do and not how it does it.

`Loppu` is built with the goal of providing exponential productivity 
increases to unique and uncommon laboratory research. It was used to build two 
unique robots during the course of the author's PhD focused on retinal gene 
delivery, both of which were used extensively to automate the author's lab 
research in multiple manuscripts that are under preparation for submission 
to biomedical journals. 

# Acknowledgements
