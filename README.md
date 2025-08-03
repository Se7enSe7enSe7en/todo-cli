# todo-cli
Todo app in the cli

# how to use
```bash
todo [OPTIONS] COMMAND PARAMS
```

# examples
## create a new todo 
```bash
todo create "feed dog"
```

### create new todo with priority
```bash
todo --priority=high create "feed dog"
``` 

## list all todos
```bash
todo list
```

| id | description            | done | created at       | priority |
---------------------------------------------------------------------
| 4  | renew driver's license | no   | 2024-11-28 20:01 | Medium   |
| 3  | ask about task QHA-968 | no   | 2024-11-28 17:38 | Medium   |
| 2  | work on todo-cli       | no   | 2024-11-27 15:32 | Low      |
| 1  | feed dog               | yes  | 2024-11-27 06:41 | High     |

### notes
- default sorting order is (done = ASC) > (priority = DESC) > (created_at = DESC)
    - sorting order, ascending = ASC, descending = DESC
    - done, yes > no, high is "yes", low is "no"
    - priority, High > Medium > Low, high is "High", low is "Low"
    - created at, newer date > older date, high is newer date, low is older date
- date is in yyyy-mm-dd HH:MM (year-month-day Hour:Minutes) format

### list with options
```bash
todo list --sort=ASC --not-done 
```

## mark todo as done
```bash
todo done 1
```