# TradeWare

## Overview

The trading package is a tool designed to assist in creating readable and clean trading strategy rules in algorithmic trading using Go. It utilizes the Composite Design Pattern and Factory Method Pattern to provide a flexible and extensible framework for building trading bots.

The package includes various predefined trading rule conditions such as crossup, crossdown, first indicator below second, first indicator above second, as well as three logic operators: And, Or, and Not. These conditions can be combined and nested to create complex trading strategies.

## Features

- Easy-to-use API for defining trading rules
- Supports multiple pre-defined trading rule conditions
- Flexible composition of conditions using logic operators
- Provides a Factory Method for creating custom rule conditions
- Encourages readable and clean code for algorithmic trading strategies

## Installation

To install the trading package, you need to have Go installed and set up on your machine. Then, you can use the following command to install the package:

```bash
go get github.com/kaialgo/tradeware
```

## Usage 
To use the trading package in your Go code, you first need to import it:
```bash
import "github.com/kaialgo/tradeware"
```
Next, you can start creating trading rules by combining various conditions and logic operators. Here's an example:
```bash
package main

import (
    "fmt"
    tw "github.com/kaialgo/tradeware"
)

func main() {
    firstIndicator := []{100, 102, 101, ... }
    secondIndicator := []{102, 104, 105, ... }
    
    // Create a rule condition for a crossup between two indicators
    crossupCondition := tw.NewCrossUp(firstIndicator, secondIndicator, 1)

    // Create a rule condition for the first indicator being below the second
    belowCondition := tw.NewFirstIndicatorBelowSecond(firstIndicator, secondIndicator, 1, 1)

    // Combine conditions using the And operator
    andCondition := tw.NewAnd(crossupCondition, belowCondition)

    // Evaluate the combined condition
    result := andCondition.isValid()

    fmt.Println("Trading rule result:", result)
}

