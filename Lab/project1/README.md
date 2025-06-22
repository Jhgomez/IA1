# Unimatch
This is a project created fully with Golang programning language that leverage the power of Prolog using Golog(Prolog for golang).

Prolog is a logic programming language used mainly for artificial intelligence and computational linguistics. Instead of giving step-by-step instructions, you define facts and rules, and Prolog uses logical inference to answer queries. It's declarative—focused on what is true, not how to compute it.

This system uses Fyne to build its UI and aims to help students find the right career match

# Environment setup
## Create Docker Sql Server Container
This will be the replacement for a local storage which would be an improvement to implement later

docker run -e "ACCEPT_EULA=Y" -e "MSSQL_SA_PASSWORD=abcdeF1+" \
-p 1433:1433 --name sql1 --hostname sql1 \
-d \
mcr.microsoft.com/mssql/server:2022-latest

## Libraries used

* golog: prolog for golang
* fyne: Golang crossplatform UI toolkit
* gin: Rest api library

## Clean Architecture
The program follows clean architecture to create testable, mantainable, scalable and flexible code. It has two applications, as Mentioned the local storage was not used because the challenges it represents that is why we have an SQL server database to store careers, aptitudes, skills and interests. We create an REST api and some sort of backend that again implements clean architecture

The frontend is a desktop application which has a data and a presentation layer, the data contains the model, storage, repositories and more and presentation layer contains features

## Golog
We use prolog to define this fact

```
career_details('faculty', 'career', ['aptitude'], ['skill'], ['interest'], 1, 1, 1)
```

we use the following matching predicates/ rules that helps us find all careers that have a match and tells us how many matched of each section
```
member(X, [X|_]).
member(X, [_|T]) :- member(X, T).

% --- count how many matches from List1 exist in List2 ---
count_matches([], _, 0).
count_matches([H|T], List, Count) :-
member(H, List) ->
(count_matches(T, List, Rest), Count is Rest + 1);
count_matches(T, List, Count).
```

Then we get all matches with this query
```
match_all_careers(%s, %s, %s, Faculty, Career, A, S, I, AT, ST, IT).
```

This returns the careers the total match and the total for each section so we can calculate percentages

# Challenges
## Create Navigation
I had to handle this and implemented a navigation stack to be able to mantain state across different screens

## Searching Rules/Predicates
It was hard to define them since Golog doesn't have utilities as Prolog