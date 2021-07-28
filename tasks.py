from invoke import run, task
import os
import re

def _get_exercise():
    cwd = os.getcwd()
    p = re.compile(".*exercises\/exercise_(\d+)")
    m = p.match(cwd)
    if not m:
        return m
    else:
        return m.group(1)

@task(
    help = {
        "exercise": "exercise to run"
    }

)
def test(ctx, exercise =_get_exercise()):
    if not exercise:
        print("ERROR: either run `inv test` within an exercise directory or pass an excerse to the invoke command. (ex. `inv test --exercise 1`)")
        return
    root_dir =  os.path.dirname(os.path.abspath(__file__)) 
    os.chdir(f"{root_dir}/exercises/exercise_{exercise}")
    print(f'Running Exercise #{exercise}...')
    print("Installing Dependencies...")
    ctx.run("go mod download")
    ctx.run("go install")
    print("Running Test...")
    ctx.run(f"go test -v excercise_{exercise}_test.go")
