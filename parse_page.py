import json
from bs4 import BeautifulSoup
import bs4
import requests


def main():

    url = "https://www.ilimgroup.ru/career/vakansii/"

    page = requests.get(url)

    print(f"{page.status_code=}")

    soup = BeautifulSoup(page.text, "html.parser")

    all_jobs = soup.find_all("article", class_="news-card-mini")

    jobs = [parse_job(job) for job in all_jobs]

    with open("jobs.json", "w") as file:
        json.dump(jobs, file)


def parse_job(job):
    job_obj = {}
    job_obj["title"] = job.find("a", class_="news-card-mini__title").text

    job_obj["tag"] = job.find("a", class_="material-footer__link").text
    job_obj["location"] = job.find("span", class_="news-card-mini__location").text

    href = job.find("a", class_="news-card-mini__title")["href"]

    job_obj["description"] = parse_description(href)

    return job_obj


def parse_description(href):

    url = "https://www.ilimgroup.ru" + href
    page = requests.get(url)

    soap = BeautifulSoup(page.text, "html.parser")

    elems = soap.find("div", class_="article__content")

    description = ""

    elems = elems.find_next("h4")
    while elems is not None:
        description += f"# {elems.text}\n"
        ul = elems.find_next("ul")
        for li in ul:
            if isinstance(li, bs4.Tag):
                description += f"* {li.text}\n"
        elems = elems.find_next("h4")

    return description


main()
