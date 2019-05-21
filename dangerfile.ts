// eslint-disable-next-line
import { message, schedule, danger } from "danger";

import { checkForNewDependencies } from "danger-plugin-yarn";
import jiraIssue from "danger-plugin-jira-issue";

function fileChanges() {
  console.log("github base: " + danger.github.base);
  return;
  const {
    modified_files: mod,
    created_files: add,
    deleted_files: del
  } = danger.git;
  console.log(Object.keys(danger.git));
  console.log({ mod });
  const title = "### File changes";

  const files = [
    ...mod.map(f => [f, "m"]),
    ...add.map(f => [f, "+"]),
    ...del.map(f => [f, "-"])
  ].sort(([str1], [str2]) => str1.localeCompare(str2));

  message(`
  ${title}
  ${"```diff"}
  ${files.map(([file, prefix]) => `${prefix} ${file}`).join("\n")}
  ${"```"}
  `);
}

// Error:  Error: Results passed to Danger JS did not include fails.
// https://github.com/danger/danger-js/issues/604
// eslint-disable-next-line
function dependencies() {
  return danger.git
    .JSONDiffForFile("package.json")
    .then(packageDiff => checkForNewDependencies(packageDiff));
}

function jira() {
  jiraIssue({
    key: ["CROVP", "PWBS", "PWC", "PWCO", "PWLP", "UP"],
    url: "https://tv2cms.atlassian.net/browse"
  });
}

function links() {
  const { number } = danger.github.pr;

  const deploymentLinks = [
    "checkout",
    "content",
    "landingpage",
    "login",
    "receiver",
    "storybook"
  ].map(site => {
    const url = `https://${site}-pr-${number}-playfe.dev.rancher.tv2net.dk`;
    return danger.utils.href(url, url);
  });

  message(
    `<ul><!-- top padding --></ul><ul><li>${deploymentLinks.join(
      "</li><li>"
    )}</li></ul>`
  );
}

// jira();
links();
fileChanges();

// schedule(dependencies());
