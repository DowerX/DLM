function handleChanged(delta) {
    if (delta.state && delta.state.current === "complete") {
      console.log(`Download ${delta.id} has completed.`);
      window.open("dlm:C:/DLM/config.yml")
    }
  }
  
  console.log("Added download manager addon.")
  chrome.downloads.onChanged.addListener(handleChanged);