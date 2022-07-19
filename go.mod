module github.com/kasrasaeed/trade_vessel

go 1.18

retract [v0.0.0, v1.0.0]
retract v0.0.0-20220719101749-1f06ba92d8b8
retract v0.0.0-20220719094859-28f4c71193cb

replace (
	github.com/kasrasaeed/trade_vessel/trade_middleware/constants v0.0.0-20220719101749-1f06ba92d8b8 => github.com/kasrasaeed/trade_vessel v1.0.1
	github.com/kasrasaeed/trade_vessel/strategy/rules v0.0.0-20220719094859-28f4c71193cb => github.com/kasrasaeed/trade_vessel v1.0.1
)