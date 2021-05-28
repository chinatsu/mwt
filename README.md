# mwt

MinWinTid but on the command line 😈

## Setup

Unfortunately, I haven't implemented a hacky exploit that grabs cookies and other necessary info from your browser of choice.
So, in order to get this running, we need a few things from MinWinTid.
The easiest way to store these things are as environment variables, so `mwt` can read them from there.

| Variable | Description |
|----------|-------------|
| `MRH_SESSION` | This is the cookie named `MRHSession` |
| `ASP_NET_SESSION_ID` | This is the cookie named `ASP.NET_SessionId` |
| `EMPLOYEE_ID` | This can be found as a field in a POST request payload, `activeIdentity.EmployeeId` |
| `POSITION_ID` | This can be found as a field in a POST request payload, `activeIdentity.PositionId` |

To get the cookies in Chrome, navigate to [MinWinTid](https://minwintid.adeo.no/MinWintid/), right click and select Inspect.
From here, select Application, then have a look under Cookies -> `https://minwintid.adeo.no`.

To get the other fields, select Network in the Inspect area and refresh to populate the requests.
There should be a request named `GetInitialData` or `GetJobAdditionSettings`, these have the POST request payload we need to look at.

To set the variables, running something like `export MRH_SESSION=MRHSession-cookie-value-here` should do the trick.
Do this for every defined field in the table above.

⚠️ `ASP.NET_SessionId` changes rather often, so things might break and you'll have to update the environment variable to get things going again.

## Running

Now that you've got the environment variables set up, you can use the Makefile actions.

- `make status` prints whether you're checked in, and your flexi time status.
- `make in` checks you in
- `make out` checks you out
