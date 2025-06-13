using UnityEngine;
using UnityEngine.UI;

public class GameDescriptionUI : MonoBehaviour
{
    [Header("UI Elements")]
    public GameObject descriptionPanel;
    public GameObject[] pages;
    public Button prevButton;
    public Button nextButton;
    public Button closeButton;
    public GameObject hintText;

    private int currentPageIndex = 0;
    private bool hasOpened = false;

    private void Start()
    {
        descriptionPanel.SetActive(false);
        ShowPage(0);

        prevButton.onClick.AddListener(ShowPreviousPage);
        nextButton.onClick.AddListener(ShowNextPage);
        closeButton.onClick.AddListener(CloseDescription);
    }

    private void Update()
    {
        if (Input.GetKeyDown(KeyCode.T))
        {
            if (!descriptionPanel.activeSelf)
                OpenDescription();
            else
                CloseDescription();
        }
    }

    void OpenDescription()
    {
        descriptionPanel.SetActive(true);
        Time.timeScale = 0f;

        if (!hasOpened && hintText != null)
        {
            hintText.SetActive(false);
            hasOpened = true;
        }
    }

    void CloseDescription()
    {
        descriptionPanel.SetActive(false);
        Time.timeScale = 1f;
    }

    void ShowPage(int index)
    {
        currentPageIndex = Mathf.Clamp(index, 0, pages.Length - 1);
        for (int i = 0; i < pages.Length; i++)
        {
            pages[i].SetActive(i == currentPageIndex);
        }

        prevButton.interactable = currentPageIndex > 0;
        nextButton.interactable = currentPageIndex < pages.Length - 1;
    }

    void ShowPreviousPage()
    {
        ShowPage(currentPageIndex - 1);
    }

    void ShowNextPage()
    {
        ShowPage(currentPageIndex + 1);
    }
}
