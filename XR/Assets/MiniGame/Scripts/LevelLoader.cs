using UnityEngine;
using UnityEngine.SceneManagement;

public class LevelLoader : MonoBehaviour
{
    public void LoadLevel1()
    {
        SceneManager.LoadScene("MiniGame/Scenes/Level1");
    }

    public void LoadLevel2()
    {
        SceneManager.LoadScene("MiniGame/Scenes/Level2");
    }

    public void LoadLevel3()
    {
        SceneManager.LoadScene("MiniGame/Scenes/Level3");
    }
}
