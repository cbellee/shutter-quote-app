export function GetQuotesFromLocalStorage(quotesName, seedData) {
    try {
        var jsonQuotes = [];
        var jsonQuotes = localStorage.getItem(quotesName);
        if (jsonQuotes === '' || jsonQuotes === null) {
            localStorage.setItem(quotesName, JSON.stringify(seedData));
            console.log("setting local storage from 'seedData'");
            jsonQuotes = localStorage.getItem(quotesName);
        }
        return JSON.parse(jsonQuotes);
    } catch (err) {
        console.log("Error getting quotes from local storage: \n" + err)
    }
}

export function UpdateQuoteToLocalStorage(quotesName, quote, seedData) {
    try {
        var quotes = GetQuotesFromLocalStorage(quotesName, seedData);
        for (var i in quotes) {
            if (quotes[i].id == quote.id) {
                quotes[i] = quote;
                localStorage.setItem(quotesName, JSON.stringify(quotes));
            }
        }
    } catch (err) {
        console.log("Error adding quote to local storage: \n" + err)
        return;
    }
}

export function AddQuoteToLocalStorage(quotesName, quote, seedData) {
    try {
        var quotes = GetQuotesFromLocalStorage(quotesName, seedData);
        var newId = GetNextAvailableId(quotes);
        quote.id = newId;
        console.log("Adding new quote with quoteId: " + newId);
        quotes.push(quote);
        localStorage.setItem(quotesName, JSON.stringify(quotes));
    } catch (err) {
        console.log("Error adding quote to local storage: \n" + err)
        return;
    }
}

export function RemoveQuoteFromLocalStorage(quotesName, quoteId) {
    try {
        var quotes = GetQuotesFromLocalStorage(quotesName);
        var itemIndex = quotes.findIndex(x => x.id === quoteId);
        console.log("Removing quoteId: " + quoteId + " at index :" + itemIndex);
        quotes.splice(itemIndex, 1);
        localStorage.setItem(quotesName, JSON.stringify(quotes));
    } catch (err) {
        console.log("Error removing quote from local storage: \n" + err)
        return;
    }
}

function GetNextAvailableId(quotes) {
    if (quotes == '') {
        return "1";
    }
    var ids = quotes.map(q => q.id);
    ids.sort((a, b) => b - a);
    var latestId = ids[0]
    latestId++;
    console.log("nextId: " + latestId);
    return latestId;
}

export function GetSuburbList(postCodes, stateName) {
    var suburbList = [];
    var p = [];

    postCodes.map(postCode => {
        if (postCode.STATE_NAME == stateName) {
            p.push(postCode.SUBURB_NAME);
        }
    });

    var sortedSuburbList = p.sort();

    for (var id = 0; id < sortedSuburbList.length; id++) {
        suburbList.push({ text: id, value: sortedSuburbList[id] });
    }
    return suburbList;
}

export function GetPostCodeFromSuburb(postCodes, suburbName) {
    let po = postCodes.filter(postcode => postcode.SUBURB_NAME === suburbName);
    if (po.length <= 0) {
        console.log(po.length);
        return null;
    }
    return po[0].POST_CODE;
}
