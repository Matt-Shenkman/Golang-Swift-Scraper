# Sports-Bet_Arbitrage-Detector

Arbitrage Detector

What is Sports Bet Abritrage?
Sports Bet Arbitrage is an event where two different betting sites, Cesar's Sportsbook and Draftkings for instance, have competing odds such that the postive money line on one of the sites is greater than the negative money line on the other. If this instance occurs it is possible to place a calcated bet on both sides (one on one team and one on the other) that guarantees a small profit return. This project aims to webscrape vegasinside.com which contains information on the odds of numerous sites to find opportunites for sports bet arbitrage.

The packages used include: colly, fmt, strings, math and strconv

There are two relevant files in this project (the rest were just learning how to web scrap in in go colly) are arbitrageDetector and arbtriageInterpretor. All of the web scraping occures in arbitrageDetector. This file outputs an array of GameData structs that include the relevant information on the details of a game. Arbitrage Interpretor takes in this information and searches for arbaitrage opportinites. It will not output just any opportunity however. Some times the odds are listed incorrectly on Vegas Insider so if an odd is set two standard deviations above the mean of the rest of the odds we assume that this is incorrect and ignore this in our final calculation. Finally the interpetor looks at all games on all avaible leagues and outputs the games that contain what we are looking for if availible.

