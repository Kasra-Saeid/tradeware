# trade middleware
This package is used as a middleware between raw data like kline historical data or ema data in a period of time And processing unit which takes action based on that raw data.
</br>
</br>
As an example:
</br>
Imagine you want to be notified when latest candle crosses up ema line.</br>
In this situation you should make a "CrossUp" rule and pass kline float64 slice, ema float64 slice and also latest bar index as 1 to that rule and pass the rule to a "Strategy" and after all call the CheckEntry on the strategy object you already created.</br>
Also note that you can add as many rules as you want to the strategy and check if all rules are satisfied any time.
</br>
</br>
I used Factory method design pattern and composite method inorder to avoid spaghetti code and reduce future maintain time.
</br>
</br>
All contributions are welcome
</br>

### Note:
full doc soon 
</br>
still in alpha

