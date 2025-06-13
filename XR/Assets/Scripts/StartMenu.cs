using UnityEngine;
using UnityEngine.SceneManagement;

public class StartMenu : MonoBehaviour
{
    public string gameSceneName;

    public void OnPlayButton()
    {
        SceneManager.LoadScene(gameSceneName);
    }
}
